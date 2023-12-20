#!/bin/bash

SRC=~/digital-datasheets-worktree1
DST=~/edatasheets.github.io

rm -r "$DST/*"

# TODO: copy README too
cp -r "$SRC/part-spec" "$SRC/support-docs" "$SRC/LICENSE" "$DST"

mv "$DST/support-docs/README.md" "$DST"

# Replace the $id line with the official release github URI in every JSON file in part-spec.
sed -i 's|"$id":.*\(/part-spec/.*\)$|"$id": "https://github.com/edatasheets/edatasheets.github.io/blob/main\1|' $(find "part-spec/" | grep '\.json$')
