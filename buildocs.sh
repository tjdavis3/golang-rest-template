TMPFILE=build.$$
echo "#!/bin/sh" > $TMPFILE
echo "pip install mkdocs-windmill" >> $TMPFILE
echo "foliant make -w mkdocs site" >> $TMPFILE
chmod +x $TMPFILE
docker run -v `pwd`:/usr/src/app -w /usr/src/app --entrypoint=/usr/src/app/$TMPFILE foliant/foliant:full 
[ -f $TMPFILE ] && rm $TMPFILE
