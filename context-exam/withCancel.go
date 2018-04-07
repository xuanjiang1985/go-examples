package main 

import (
	"log"
	"time"
	"golang.org/x/net/context"
)

func main(){
	timeoutHander()
}


func timeoutHander() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	//go doStuff(ctx)
	go doTimeOutStuff(ctx)
	time.Sleep(10 * time.Second)

	cancel()
}

func doStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            log.Println("done")
            return
        default:
            log.Println("work")
        }
    }
}

func doTimeOutStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)

        if deadline, ok := ctx.Deadline(); ok { //设置了deadl
            log.Println("deadline set")
            if time.Now().After(deadline) {
                log.Println(ctx.Err().Error())
                return
            }

        }

        select {
        case <-ctx.Done():
            log.Println("done")
            return
        default:
            log.Println("work")
        }
    }
}
