[ -d .diagramscache ] || mkdir .diagramscache
TMPFILE=build.$$
echo "#!/bin/sh" > $TMPFILE
echo "cp spec/openapi.yaml docs" >> $TMPFILE
echo "pip install mkdocs-windmill mkdocs-awesome-pages-plugin " >> $TMPFILE
echo "foliant make -w mkdocs site" >> $TMPFILE
chmod +x $TMPFILE
docker run --rm -v `pwd`:/usr/src/app -w /usr/src/app --entrypoint=/usr/src/app/$TMPFILE foliant/foliant:full 
[ -f $TMPFILE ] && rm $TMPFILE
