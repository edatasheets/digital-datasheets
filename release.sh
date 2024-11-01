#!/bin/bash

# Run the readmegenscript


SRC=~/Documents/digital-datasheets
DST=~/Documents/edatasheets.github.io

rm -r "$DST/*"

# TODO: copy README too
cp -r "$SRC/part-spec" "$SRC/support-docs" "$SRC/examples" "$SRC/LICENSE" "$DST"

mv "$DST/support-docs/README.md" "$DST"
rm -r "$DST/support-docs"

# Rename the source information in the README with the official release github URI
sed -i 's|https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/|https://github.com/edatasheets/edatasheets.github.io/blob/main/part-spec/|g' "$DST/README.md"
sed -i 's|https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/|https://github.com/edatasheets/edatasheets.github.io/blob/main/part-spec/|g' "$DST/README.md"

# Replace the $id line with the official release github URI in every JSON file in part-spec.
sed -i 's|"$id":.*\(/part-spec/.*\)$|"$id": "https://github.com/edatasheets/edatasheets.github.io/blob/main\1|' $(find "part-spec/" | grep '\.json$')
