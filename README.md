# Use the CLI app

---

## Make sure you have go downloaded

---

### To run the CLI

run the program with arguments:</br>
1 argument is encrypt or decrypt, </br>
2 argument is file path </br>
3 with 'your key' key must be 16, 24 ,32 characters </br>

In your terminal run to encrypt

```zsh
go run . \"encrypt\" \"exampleFile.txt\" \"Example key\"
```

to decrypt

```zsh
go run . \"decrypt\" \"exampleFile.txt\" \"Example key\"
```
