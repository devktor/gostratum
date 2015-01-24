package gostratum

import(
    "sync"
    "time"
    "errors"
)

type Action struct{
    cond  *sync.Cond
    response *Response
    active bool
    timer *time.Timer
}


func MakeAction() *Action{
    return &Action{
        cond: &sync.Cond{L: &sync.Mutex{}},
        response: nil,
        active: true,
        timer: nil};
}

func (a *Action) SetTimeout(timeout time.Duration, callback func()){
    a.cond.L.Lock()
    defer a.cond.L.Unlock()
    if a.timer != nil {a.timer.Stop()}
    if timeout != 0 {
        a.timer = time.AfterFunc(timeout, func(){
            callback()
            a.Done(&Response{Error:errors.New("timeout")})
        })
    }else{
        a.timer = nil
    }
    
}

func (a *Action) Wait() *Response{
    a.cond.L.Lock()
    if a.active{
        a.cond.Wait()
    }
    a.cond.L.Unlock()
    return a.response
}


func (a *Action) Done(response *Response){
    a.cond.L.Lock()
    a.response = response
    a.active = false
    if a.timer != nil {a.timer.Stop()}
    a.cond.L.Unlock()
    a.cond.Signal()
}
