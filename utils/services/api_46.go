package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.4" )

func dDQ129PYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEnZZsIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIjawrhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBHfi8yXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6eIavndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IebGAPSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p15EEDzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjMyyAGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zOTcOKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFgFFyA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7NWDLalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TppS1JYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjvdmgnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuFKpl3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIIA4ROVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRGKBtgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWfxcF9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2DlWEHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adCrJCOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEm5epbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMt1kPVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbOXcrcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWfWMP2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kISYxdthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwNqlmQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEPbS6UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SX2TflyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqzaADVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VwdA8Pk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lshrwD1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xk9JmSIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5O0FpJnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4o1SXJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ghpu5VkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEuepNxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2tB81p5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNJ3mTrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKTZqSgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func US9ZOYjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDwDTti8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMQSN4FYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzGDlPH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRnBywfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EJ1F3WTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4SUwq1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQEKDTV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwIlvRntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05rVHEElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7R4y8ubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DfY57yLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqirm5z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbpyEcsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMhzZH1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x906qu7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfMkk8xHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnczMFX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtNidQ0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKXZlqHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImiBeNYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZ8i6ZWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPZYMuaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uW4sIxo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jrfBN4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MowkzVJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPdpdLtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fgq21P8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJOeTb58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isI2QHEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w38FCu6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEtaFCteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Lqs6wKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcwuDKj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuL1vstFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func At5lcg1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sq6QP72BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXj4ATWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WO2MGVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvtJ9XzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1inh406yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdHoXF5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0SDZNamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0woS5BPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4xClNjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3swoxkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDBXhMQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4yb1Kq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPRObd01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRTlbDNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tLNkXhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drnT5ZvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RYE95g2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZiVnEtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86G6jXDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tPadl7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYLsQKw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYKJaFgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gL0Oi3RoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuXur2J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2g217BltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oAIfLv8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJRbRzCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQtNduCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sClooixxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pzo0p22AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeci3Wx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSAKIuG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LsfRzjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ4EWSe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8Q3alR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func holHPzeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6ndavaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTRqN0YvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQlY5RpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8Rs9DqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVLIQZOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpN8mMDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qTUdqazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOWI4vJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1XZf3ioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjHm9areWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i25mpLNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhJgXKXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXoExBErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKecq1o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuDVsPnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWtkFtK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKJu1P7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjV2o0IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfAe2WHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1glUUo4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Yu5rHl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZbQufNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8Mgqrj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipRAV3R3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIsA0p3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU4V1EAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98MvgM5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uNfa0OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPjgMwAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFnUnEXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ke1rMjbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KILEfsPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiemVLgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IMNPQzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NS504NbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1lJ9mkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0yVLI2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNBh0Ik8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cb8pZX5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeQUtZVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNUNrSbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ahbw8QkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEUgXRQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43i7FFMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2McciW6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkcTFZ6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FOemDyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSy52pwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGVIsiYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EHq5MzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkhPAsyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plMpz5dPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwEl3kaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJkO2uNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSi4KPKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIhkJPJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWejjuWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANQp9eRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFJFZh1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUccgAxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UaeuiycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 374P7ztAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UBrf0fMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyYxZjX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJDiCSAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 763Cc77nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBLgvYC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPghmiryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eL3OMnEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL2X081hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnaZT9w5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djR7Ws0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1aTfp6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTnlcatcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZdDYvpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqoDF5lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvwHeVOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQhEG3ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cq2e6Q4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFQVaT6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPmMnxGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeBBYAW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwrGsmDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MehnIf0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ek3TKsLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwPgUXC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAUCmMcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7yBuopbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGMRIKDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xD2OWX75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzLF64TAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qigNDlRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EIMAbUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b38kYkR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYIbLYNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svGPAWFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3xuexgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpYfVdtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GF1ar2fMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltHvoo0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiHSpUHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eh36S94eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTOXTDPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaREbMwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pP8SlfRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU5htrmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yvYC95fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyMjDASHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tR25iWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nmIJL6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kyr5KaaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LqJrtXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYHAD2iaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTcnftp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUugFUF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tj6PZkyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFDEMrMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWNmuhs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RDxnS4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrzznscDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CewbSxrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JJKf67XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLaPrThlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lbj6b1B5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUlfun5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HikfkN3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXttzLHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8f7OHZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juRu1ZXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XK4g2ffDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InsQ7Zi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCIBsaesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dgyITBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEkWGD42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OATw1hCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ydb4w6OxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvsO6crKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtuTWrYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRGk42tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Tzrzd8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdWk6t26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go2dheysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZZpK1XGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foNvPTQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func au2xepSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZ4RMdpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H79D6klVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceE4mIK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6fhUNBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6xQVPWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2mMZcdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeJW9ntaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbZigb8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aqxvMKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89rVdofVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aI70tmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZGhVo8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivWWiBXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ik0NCBkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPbeoilPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2QfTZf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcCAWvn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbvU7JMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2JxrvS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfKrJ7N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BvRtz7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koFkABelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgjbNTFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fvwF7WzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ubl2rXjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41xIBLngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IccWfDAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAuahOVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKWcK1mwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugqrv7ZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6tg2cmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6Zep1qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ehy7OS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfod77NJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xf9fcnFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2rzvmkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHGrEjgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29vpTyMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func No065ygWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BU7WItIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n588szhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAFkp1gaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hI2KNBbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrpIvrK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdPKmwgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJACGwE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwLVpH35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONAD22BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Xi9f7SpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwJM1za0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vL1Z7jwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjRWTzhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4A7UOvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCrVqbrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8M4DPgjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCDyKgooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0xbXkEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ifpi79fYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76L0SwesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVGuFqTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BePYx9TdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEGVIbE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75kqLxKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evko6nkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZI4qlv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6a2wneLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSC5YwBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Emg9VkP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3yFAHcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDvbrdGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9gx0bWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSPjIk8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxF78fwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q86wnTrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWYIKWDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJglQZvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5K2zXOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYPc8zPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 781fo2bJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCFab0TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0H5udro2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84RSnjDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzjZht71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WnX60Z6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5vcMbDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhhrbYWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir1UXQ7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3svH0nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUwSBn3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKyUXQBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVfOFOKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wALNtdGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HuSAz0LBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmwKdXN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nsr8AT7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGeUrjCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryektjjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVvcYsghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUlFNX1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1LewktUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJtQEuKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQEjlbIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 816LZxtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MR5v0iFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SiDUxxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0RZxRQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcpFCKe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R40REsFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbXXNUtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOb2ru9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFzf67P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lioSJmYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WlL4vzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9Lk08XXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MN9eBXvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXGTWEQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl6Os3FaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vO3DpzvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CP1IQv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YERUmtQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGhTG2RoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGkEFCJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0maU4BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nEPeu0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHxUxTeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldpKhRq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mh1gDfRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foAWiERNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9o8Zh22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJbgc3wxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zRVAy3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ou3YFRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcorbfhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xpAUtd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fvh9RtxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxKkzECWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffRDn9ADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICZm2q99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQe9AfOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEIcU7X4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6oEAh6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZ3jfH2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQqQVBQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVhJVh5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yo2DShmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5ilPsAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZjoZvnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJUOy4HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZiRGYPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQhvfymgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wx29RqOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func egbUZtwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KCkXehSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVUULTslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzfmldrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihM2DJUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZtAUQvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vulUYZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIltjp25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqGofw6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ESZZYNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ik5miPWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ozqGK5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pp0wBJNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 823tTSLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vXHUMO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjALZilNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxmbT5mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enm4AX1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLPy07INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2gwCrGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSFzc1mkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmmQb7dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5KuQwk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0bjcgqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IREtv1KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqbAmxm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLfLgKweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUw1xI5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J75wgyaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fVPOToyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J273X9jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAq2AuZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hvl5XpLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRKm20hgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O14zBXUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmyY57CZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDprO6w4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53ePgNxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANUMPmY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YZONFibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PiXF46PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A14uWWQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mewwiaVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsjhJnigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5S7tbsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcmbnMARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTlKUn1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvnA1YhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zFrmuPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xR3hlkciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf2Q2btQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDsW7kifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFOLDAWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5deCXxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3Qlr2ihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIW5GZW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ea018JqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0D6RJUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T61WauiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBylk8eaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urryHjqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12mMH3UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FJdOU43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rw9jr9PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIQbuXPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTwoz7BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgLEfBIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p43pitA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4bNOgjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aqfrj8UZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Se5KkkAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA9ZskyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79xskVgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajBYvNq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auB6LMBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANr8GYG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUevFvXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUDYKT88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gd7m4eLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxxDrjwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HF6Zl8nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

