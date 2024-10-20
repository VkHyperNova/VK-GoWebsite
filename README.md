# VK-GoWebsite

node.js
tailwind

// Compile output.css
npm run build:css

// Auto compiles output.css
npx tailwindcss -i ./static/styles.css -o ./static/tailwind/output.css --watch

// Run prettier plugin
npx prettier --write .

or add "format": "prettier --write ." to package.json

npm run format
