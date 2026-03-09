#### 🏝️ Palm
... is tiny, simple, ease-to-use building system

#### 💡 Example
```
build {
  run "cargo build --release"
}

commit(mention) {
  print "performing commit..."
  run "git commit -m " + mention
  print "done"
}

push {
  run "git push -u origin main"
}

upd(mention) {
  do commit(mention)
  do push
}

test {
  mention = "Hello, world!"
  do commit(mention)
  do push
}
```
