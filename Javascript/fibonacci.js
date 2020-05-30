function fibonacci(b) {
  // eslint-disable-next-line
  if(b == 0){
    return 0
  }
  // eslint-disable-next-line
  if(b == 1){
    return 1
  }
  return fibonacci(b-1) + fibonacci(b-2)
}

console.log(fibonacci(6))
