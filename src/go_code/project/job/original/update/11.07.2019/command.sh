#!/bin/bash

function do_job(){

        for file in `ls $1`
        do
            if [ -d "$file" ] 
            then 
                #echo "$file is directory"
                cat "./"$file/commands.txt | ./stream.sh  | asciinema rec --title="Scone ${file}" "./"${file}"/"${file} 2>&1 | tee -a my.log
asciinema upload "./$file/"${file} | echo "${file} <$(grep "\/a\/" | sed 's/^[ \t]*//g')>" >> log.txt
            else
               # echo "$file is file"
                continue
            fi
        done


}
do_job $1