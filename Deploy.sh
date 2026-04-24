#!/bin/bash

echo "🔧 Building Sphinx documentation..."
cd docs
make html

echo "📦 Copying build to root /website..."
rm -rf ../website
mkdir ../website
cp -r _build/html/* ../website/

echo "🚀 Deployment ready. Commit and push to GitHub Pages branch."
