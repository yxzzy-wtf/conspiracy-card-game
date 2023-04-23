#! /bin/bash

# Setup tmp folder
mkdir tmp

# Compile files & card items (incl. any args)
cp template/cards.tex tmp/cards.tex
go run main.go "$@" >> tmp/cards.tex
echo "\end{document}" >> tmp/cards.tex

# Compile!
pdflatex tmp/cards.tex

mv cards.pdf latest.pdf

if [ "$1" == "-snap" ]; then
    cp tmp/cards.tex "snapshot/$(date +"%Y%m%d-%Hh%M").tex"
fi

rm cards.aux cards.log
rm tmp/*
rmdir tmp