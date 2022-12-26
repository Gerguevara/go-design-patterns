// condicion de carrera es un problema que ocurre cuando varias runtinas estas manejando 
y accediendo al mismo valor, dado q una puede modificar el valor y otra go runtina ni enterarse y seguir trabajando
con ese valor viejo y asi muchas veces, estos es un problema pq en cierta forma muchas operaciones se anularian


Al compilar GO nos da un waringin al hacer la compilacion de que esto puede ocurrir

para ello utilizamos un elmento del paquete sync (el mismo de los waitgrouds) llamado mutex

`	var lock sync.Mutex `

basicamente lo que ahces es decirle a la applicacion que se bloquee por completo utilizando el metodo 

`lock.Lock()` entonces nina optra rutina podra escribir hasta que esta termine de escribir

luego para desbloquerar utilizamos el metodo 

`	defer lock.Unlock()`
usar deffer no es necesario

Basicamente esto ahce que la go runtina que llegue primero va a bloquear a la que llegue despues hasta q la 1ra 
haga su operacion y remueva el candado