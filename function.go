func sumNumbersFromFile(filename string) int {
    // Здесь обычно идет код, который загружает содержимое файла filename
    // Например, используя ioutil.ReadFile(filename) или другие способы
    
    // Для примера, создадим слайс чисел для имитации данных из файла
    numbers := []string{"1", "2", "3", "4", "5"}

    sum := 0
    for _, numStr := range numbers {
        num, err := strconv.Atoi(numStr)
        if err != nil {
            fmt.Println("Ошибка преобразования числа:", err)
            continue
        }
        sum += num
    }

    return sum
}
// Change1