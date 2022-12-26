El problma que se plantea aca es que paso si al mismo tiempo se estan ejecutano operacions de lectura y escritura
si utlizaramosn el metodo lock de mutex  `lock.Lock()`, este bloqueria innecesariamente la lectura,
Lo que Go nos ofrece es un mecanismo para tener `1 escritor y muchos lectors`

para ellos se hace uso de 	`defer lock.RUnlock()` y 	`lock.RLock()`, si utlizaramos el mismo  `lock.Lock()` practicamente
solo tocaria esperar a que una operacion de escritura termine.


.RUnlock() nos permite utilizar un lock diferente para hacer las lecturas