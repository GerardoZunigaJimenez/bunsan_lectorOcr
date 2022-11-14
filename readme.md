1- Download the project in your local

2- Run 
    
    go mod vendor

3- The way to test this project is using a command in the consoele like this 

    go run main.go read --inputFilePath files/testFiles/entradas.txt

4- For some reason the go mod vendor is not downloading all the time one library, execute it manually
IE: 
    
    go mod download github.com/urfave/cli

and then run step 3, again

5- Custom Path for the Project, you need to change the inputFilePath value

    go run main.go read --inputFilePath {inputFilePath}

The idea is to have a simple interface like input data to help with the consume
