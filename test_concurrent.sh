#!/bin/bash


if [[ ! -d "test_results" ]]; then
   mkdir "test_results"
else 
    rm "test_results/*"
fi



for n in {0..9}
do
    time php -f test_server.php "tmp$n.jpg" &
done