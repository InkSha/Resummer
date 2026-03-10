#!/bin/sh

echo "Build Application..."

cd mobile

echo "Build Mobile App"

flutter build apk

cd ..

echo ""

if [ ! -d "frontend/public/apps" ]; then
  mkdir frontend/public/apps
fi

if [ -f "frontend/public/apps/app.apk" ]; then
  rm frontend/public/apps/app.apk
fi

mv mobile/build/app/outputs/apk/release/app-release.apk frontend/public/apps/app.apk

echo ""

cd frontend

echo "Build Frontend App"

pnpm run build

cd ..

if [ -d "dist/out/" ]; then
  rm -rf dist/out/
fi

if [ ! -d "dist" ]; then
  mkdir dist
fi

mv frontend/out dist/out

echo ""

cd server

echo "Build Server App"

go build -o ../dist/server.exe .

cd ..

if [ -f "dist.tar.gz" ]; then
  rm dist.tar.gz
fi

tar -czvf dist.tar.gz dist/

clean() {
  if [ -d "dist/" ]; then
    rm -rf dist
  fi

  if [ -f "frontend/public/apps/app.apk" ]; then
    rm frontend/public/apps/app.apk
  fi
}

clean
