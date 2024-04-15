package stack

import (
	"math"
)

type Node struct {
	Value interface{}
	Prev  *Node
}

type Stack struct {
	Length uint
	Head   *Node
}

func New() Stack {
	return Stack{Length: 0, Head: nil}
}

// Metemos un nuevo elemento al stack, por lo que ahora la cabeza de nuestro stack va a ser el ultimo elemento que metimos

func (s *Stack) Push(item interface{}) {

	// Aumentamos la Length de nuestro stack porque vamos a ingresarle un nuevo elemento

	s.Length++

	// Creamos el nodo del nuevo elemento y de momento su propiedad de prev no la enlazamos con nada

	node := Node{Value: item, Prev: nil}

	if s.Head != nil {
		// Si el head no es nulo, es decir, si existe al menos un elemento en nuestro stack, hacemos que el prev del nuevo nodo sea el head actual del stack
		node.Prev = s.Head
	}

	// Actualizamos el head del stack para que ahora apunte al nuevo nodo que ya  se encuentra enlazado con el anterior head
	s.Head = &node

}

func (s *Stack) Pop() interface{} {
	s.Length = uint(math.Max(0, float64(s.Length)-1))
	// Aqui lo que hacemos es que retornamos el numero mas grande entre dos numeros

	// Si nuestra Length - 1 es menor a 0, la Length se mantendra en 0, pero si no es asi, es decir nuestra lenght menos uno es mayor a 0 pues el valor de nuestra length se actualiara a ese

	// Esto sirve para evitar tener que poner un if de que si nuestra lenght es 0 no se haga la operación

	// if (s.Length != 0) {
	//     s.Length--
	// }

	// Guardamos la head actual en una variable
	oldHead := s.Head

	//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
	//                                   ^
	//                                   head

	// Si la Length despues de actualizarla es 0, signfica que vamos a vaciar nuestro stack, poir lo que hacemos que simplemente la head este apuntando a nil
	if s.Length == 0 {
		// (A)
		// ^
		// head

		// nil   (A)
		// ^
		// head
		s.Head = nil
	}

	// Si nuestra head no es nil, signfica que podemos sacar elementos de nuestro stack porque tiene elementos, por lo que, hacemos que nuestro head ahora apunte al elemento previo al head antiguo, y desenlazamos al head antiguo de nuestro stack
	if s.Head != nil {
		// Estado actual de nuestro stack
		//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
		//                                   ^
		//                                   head
		s.Head = oldHead.Prev
		// Hacemos que head apunte ahora al elemento previo al que se le va a hacer pop
		//(A) <- (B) <- (C) <- (D) <- (E) <- (F)
		//                            ^
		//                            head
		oldHead.Prev = nil
		// Desenlazamos el ultimo nodo (el que queremos sacar)
		//(A) <- (B) <- (C) <- (D) <- (E)    (F)
		//                            ^
		//                            head
	}

	return oldHead.Value
	// Devolvemos el valor del elemento que sacamos
}

// Retornamos el valor contenido en el head de nuestro stack
func (s *Stack) Peek() interface{} {
	return s.Head.Value
}
