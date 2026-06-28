#!/bin/sh
set -e

for file in /usr/share/nginx/html/assets/*.js; do
  [ -f "$file" ] && sed -i "s|__VITE_IDENTITY_API_URL__|${VITE_IDENTITY_API_URL:-http://localhost:3334}|g" "$file"
  [ -f "$file" ] && sed -i "s|__VITE_SOCIAL_API_URL__|${VITE_SOCIAL_API_URL:-http://localhost:3335}|g" "$file"
done

exec "$@"
