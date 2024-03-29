package lua

import(
    "sync"
)

// GlobalRegistry holds golang values in which their registry indexes are to be passed 
// to the C runtime. upon return from the C runtime, we can use the Registry Indexes 
// sent back from C to obtain our real values. NEVER EVER SEND GOLANG POINTERS TO C. 
// golang may change pointer values during scheduling and GC, so references to go 
// pointers in C may age-out
var GlobalRegistry *Registry = NewRegistry()

// GlobalMutex provides a mutex across the entire package.
var GlobalMutex *sync.Mutex = &sync.Mutex{}