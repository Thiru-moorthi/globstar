// <expect-error> File not closed within function
func badFileHandling() error {
    f, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    data := make([]byte, 100)
    _, err = f.Read(data)
    if err != nil {
        return err  
    }
    return nil
}