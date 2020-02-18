workflow "Unit Tests" {
  resolves = ["cedrickring/golang-action@1.4.1"]
  on = "push"
}

action "cedrickring/golang-action@1.4.1" {
  uses = "cedrickring/golang-action@1.4.1"
}
