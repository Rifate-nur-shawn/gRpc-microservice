package sample

import (
    "math/rand"
    pb "gRpc-microservice/pb"
)

func randomKeyboardLayout() string {
    layouts := []string{"QWERTY", "QWERTZ", "AZERTY"}
    return layouts[rand.Intn(len(layouts))]
}

func randomBool() bool {
    return rand.Intn(2) == 1
}

func randomCPUBrand() string{
    return randomStringFromSet("Intel","AMD")
}

func randomCPUName(brand string) string {
    if brand == "Intel" {
        return randomStringFromSet("Core i3", "Core i5", "Core i7", "Core i9")
    }
    return randomStringFromSet("Ryzen 3", "Ryzen 5", "Ryzen 7", "Ryzen 9")
}

func randomInt(min, max int) int {
    return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func randomStringFromSet(a ...string) string{
    n:=len(a)
    if n ==0 {
        return ""
    }
    return a[rand.Intn(n)]
}