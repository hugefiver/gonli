for i in *
do
	if [[ $i =~ (packs|shell.sh) ]];
	then
		continue
	fi

	zip -r -9 packs/${i}.zip $i
done

