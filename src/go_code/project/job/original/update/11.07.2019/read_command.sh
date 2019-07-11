#!/bin/bash

function do_job(){
        for file in `ls $1`
        do
            MaxTime=1.0
            if [ -d "$file" ] 
            then 
                #echo "$file is directory"
                for file1 in `ls ./$file`
                do
                    if [ "$file" = "$file1" ]
                    then
                        count=0
                        T_last=""
                        T_abs=""
                        #statements
                        cat "./"$file/$file | while IFS='' read T_current
                            do
                                ((count++))
                                if [ "$T_last" = "" ]
                                then
                                    T_last=$T_current
                                    T_abs=$T_current                                    
                                    continue
                                else
                                    #tmp=`echo $line |  grep -Po '(?<=(\[ )).*(?= ,)'`
                                    #get timestamp

                                    tmp1=${T_last%%, *}
                                    tmp1=${tmp1##*[}

                                    tmp2=${T_current%%, *}
                                    tmp2=${tmp2##*[}
                                
                                    tmp3=${T_abs%%, *}
                                    tmp3=${tmp3##*[}
                                    #avoid selecting the first line 
                                    case "$tmp1" in
                                      [0-9]*)

                                        #float calculation
                                        Delta=`echo "$tmp2-$tmp1"|bc`
                                        if [ $(echo "$Delta > $MaxTime"|bc) -eq 1 ]
                                        then
                                            Delta=$MaxTime
                                        fi
                                        add=$(printf "%.6f" $(echo $tmp3+$Delta|bc))
                                        sed -i '.bak' "${count}s/[[0-9.]*/[${add}/" ./$file/$file
                                        T_abs=`echo $T_abs| sed "s/[[0-9.]*/[${add}/"` 
                                        T_last=$T_current
                                    ;;
                                     *)
                                     T_last=$T_current
                                     T_abs=$T_last
                                     ;;
                                    esac
                                fi
                            done

                    fi
                done
            else
               
                continue
            fi
        done


}
do_job $1