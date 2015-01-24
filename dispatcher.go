package gostratum

import (
    "sync"
    "errors"
    "encoding/json"
)

type Dispatcher struct{
    pending map[uint64] *Action
    callbacks map[string] func(*json.RawMessage)
    mutex *sync.Mutex
}

func MakeDispatcher() *Dispatcher{
    return &Dispatcher{pending:make(map[uint64] *Action), callbacks: make(map[string] func(*json.RawMessage)), mutex: &sync.Mutex{}}
}


func (d *Dispatcher) RegisterRequest(request *Request) (*Action, error){
    action:=MakeAction()
    d.mutex.Lock()
    defer d.mutex.Unlock()

    if d.pending[request.ID] !=nil{
        return nil,errors.New("id already taken")
    }
    d.pending[request.ID] = action
    return action,nil
}

func (d *Dispatcher) RegisterNotifiactionHandler(uri string, callback func(*json.RawMessage)){
    d.mutex.Lock()
    defer d.mutex.Unlock()
    d.callbacks[uri] = callback
}

func (d *Dispatcher) Cancel(id uint64){
    d.mutex.Lock()
    defer d.mutex.Unlock()
    delete(d.pending, id)
}

func (d *Dispatcher) Process(msg *Response){
    d.mutex.Lock()
    action := d.pending[msg.ID]
    if action!=nil{
        delete(d.pending, msg.ID)
        d.mutex.Unlock()
        action.Done(msg)
    }else{
        callback:=d.callbacks[msg.Method]
        d.mutex.Unlock()
        if(callback!=nil){
            callback(msg.Params)
        }
    }
}

func (d *Dispatcher) Error(err error){
    response := &Response{Error:err}
    d.mutex.Lock()
    defer d.mutex.Unlock()
    for _,action := range d.pending{
        action.Done(response)
    }
}

