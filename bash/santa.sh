#!/bin/bash


actionSanta ()
{ # what should we do with our santa list?
	echo $1 "buys for" $2
	#If you want to mail rather than displaying:
	echo "Hi $1, you buy for $2" | mail -s "Secret Santa " $1
} 

#Shuffle and write to new file
shuf emails.txt > /tmp/shuffnames.txt

#Read into memory
readarray shuffednames < /tmp/shuffnames.txt

firstPerson=${shuffednames[0]}
count=0
for i in "${shuffednames[@]}"
do
	if [ $count -gt 0 ]; then
		actionSanta $i $prevName
	fi
	prevName=$i
	count+=1
done

actionSanta $firstPerson $i
