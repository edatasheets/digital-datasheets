#!/bin/bash

SRC=~/Documents/digital-datasheets
DST=~/Documents/edatasheets.github.io

rm -r "$DST/*"

# TODO: copy README too
cp -r "$SRC/part-spec" "$SRC/support-docs" "$SRC/examples" "$SRC/LICENSE" "$DST"

mv "$DST/support-docs/README.md" "$DST"
rm -r "$DST/support-docs"

cd "$DST"


# Replace the $id line with the official release github URI in every JSON file in part-spec.
sed -i 's|"$id":.*\(/part-spec/.*\)$|"$id": "https://github.com/edatasheets/edatasheets.github.io/blob/main\1|' $(find "part-spec/" | grep '\.json$')

# Fix readme links
sed -i 's|https://github.com/edatasheets/digital-datasheets/blob/main|https://github.com/edatasheets/edatasheets.github.io/blob/main|g' README.md
