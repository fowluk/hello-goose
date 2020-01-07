#!/bin/sh

SCRIPT_DIR=$(dirname $0)
sed -e "s/%CF_INSTANCE_GUID%/$CF_INSTANCE_GUID/g" \
    -e "s/%CF_INSTANCE_INDEX%/$CF_INSTANCE_INDEX/g" < $SCRIPT_DIR/index.html.template > $SCRIPT_DIR/public/index.html

$HOME/boot.sh
