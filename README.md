#### 🏝️ Palm
... is tiny, simple, ease-to-use building system

#### 💡 Example
```
build {
  run "cargo build --release"
}

commit (mention) {
  echo "performing commit..."
  run "git commit -m" <> mention
  echo "done"
}

push {
  run "git push -u origin main"
}

upd (mention) {
  do commit(mention)
  do push
}

test {
  let mention "Hello, world!"
  do commit(mention)
  do push
}
```
