recreation project for problem solving practice

`go test -v . | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''`

`nodemon -q --exec go test . -e go --ignore '*_test.go'  --signal SIGTERM`
