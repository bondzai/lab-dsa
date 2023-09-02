const formatinputNumberber = (inputNumber) => {
    return inputNumber.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1,')
}