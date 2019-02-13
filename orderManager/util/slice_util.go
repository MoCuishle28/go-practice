package util


// 删除字符串切片的零值
func RemoveZero(slice []string) []string{
    if len(slice) == 0 {
        return slice
    }
    for i, v := range slice {
        if v == "" {
            slice = append(slice[:i], slice[i+1:]...)
            return RemoveZero(slice)
            break
        }
    }
    return slice
}