#!/bin/bash

# Calculate the SHA256 hash of a file and return the first 6 characters
function calculate_hash() {
    local filename="$1"
    local hash=$(shasum -a 256 "$filename" | cut -c1-6)
    echo "$hash"
}

echo "Running the public hasher"

# Remove the old CSS file
rm ./public/output.*.css

# Generate the new CSS file
tailwindcss --minify -i ./styles/styles.css -o ./public/output.css

# Add the 6 first characters of the hash to the file name
# of the generated CSS file. This is to bust the cache of the CSS file.

# Get the hash of the generated CSS file
hash=$(calculate_hash ./public/output.css)

# Rename the generated CSS file
mv ./public/output.css ./public/output.$hash.css

# Replace the old CSS file with the new one
# The first argument is an empty string to skip backup in macOS
# Uses a counted range, {0,1}, to simulate a ? operator for the hash
# This matches both styles.css and styles.hash.css
sed -i "" "s/output\(\.[a-z0-9]\{6\}\)\{0,1\}\.css/output\.$hash\.css/g" ./views/layout.templ

echo "Tailwind generated and updated"
