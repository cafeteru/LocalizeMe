#! /bin/bash
# converts all puml files to svg

BASEDIR=$(dirname $(dirname $(realpath "$0")))/assets
rm "$BASEDIR"/dist -R
mkdir -p "$BASEDIR"/dist
for FILE in $BASEDIR/*.puml; do
  echo Converting "$FILE"..
  FILE_SVG=${FILE//puml/svg}
  cat "$FILE" | docker run --rm -i think/plantuml > "$FILE_SVG"
  docker run --rm -v "$PWD":/diagrams productionwentdown/ubuntu-inkscape inkscape /diagrams/"$FILE_SVG" --export-area-page --without-gui &> /dev/null
done
mv "$BASEDIR"/*.svg $BASEDIR/dist/
echo Done