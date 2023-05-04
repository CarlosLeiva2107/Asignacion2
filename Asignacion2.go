//Asignacion #2
//Stacy Chacón Argüello\2021022405
//Carlos Leiva Medaglia\2021032973
//Funciones desarrolladas de la 1 a la 10

package main

import (
	"fmt"
	"math"
)

//#########################################FUNCIONES DESARROLLADAS############################################

// 1. Función que genera una secuencia de tamaño n con números pseudo-aleatorios.
func generarPseudoAleatorios(n int, s0 int) []int {

	//Se revisa que la semilla y el n esten en el rango dado
	if s0 < 11 || s0 > 257 {
		return []int{} //De no ser asi, se retorna un arreglo vacio ya que no se cumplio con el rango
	} else {
		//Se verifica que sea primo
		for !esPrimo(s0) {
			s0++ //Mientras no cumpla que es primo, se aumenta en 1
			//Asi se obtendrá el siguiente numero primo mayor
		}
	}

	//Definir valores para a, m y b (Utilizando valores que se usaron en el libro de Dromey)
	m := 4096
	b := 853
	a := 109

	//Definir Slice donde guardaremos los números
	numeros := []int{}
	//Se usa un ciclo for para generar la secuencia de números del tamaño deseado
	for i := 1; i <= n; i++ {
		s0 = (a*s0 + b) % m //Se usa la formula para el método de congruencia lineal multiplicativa

		//Se verifica que el numero generado este en el rango valido
		//De esta manera se evitan problemas, como números negativos
		if 0 <= s0 && s0 <= (m-1) {
			numeros = append(numeros, (s0 % 100)) //Se hcae modulo 100 para que este en el rango de 0 a 99
		}
	}
	return numeros
}

// Función Complementaria, para revisar si un numero es primo o no.
func esPrimo(s int) bool {
	for i := 2; i < s; i++ {
		if s%i == 0 {
			return false
		}
	}
	return true
}

// 2. Función que genera un gráfico de barras a partir de un arreglo de tamaño n.
// (Falta, no se si es generarlo mediante libreria en una interfaz o en consola)
func graficoBarras(arreglo []int) {
}

// 3. Función que hace la inserción de un valor entero en un arreglo de numeros enteros.
// Se hace la inserción en caso de que no este en el arreglo, y esta se hace al final del arreglo.
// Se devuelve la cantidad de comparaciones realizadas tanto en la búsqueda como en la comparación.
// La función recibirá el arreglo como un puntero para que este pueda ser modificado dentro.
func insertarNumeros(arreglo *[]int, llave int) int {

	comparaciones := 0 //Para llevar registro de la cantidad de comparaciones
	//Se revisa cada dato en el arreglo
	for i := range *arreglo {
		comparaciones++
		//En caso de que exista la llave, no se modifica y se retorna la cantidad de comparaciones hechas
		if (*arreglo)[i] == llave {
			return comparaciones
		}
	}

	//En caso de no encontrarse, se agrega la llave al arreglo
	// La inserción se cuenta como una comparación
	comparaciones++
	*arreglo = append(*arreglo, llave)

	return comparaciones
}

// 4. Función que ordene ascendentemente un arreglo de enteros, mediante ordenamiento de selección
func ordenarSeleccion(arreglo []int) []int {
	//Declarar variables para usar en el ordenamiento
	min := 0 //Guarda el valor mínimo del array
	//Guarda las posiciones que van a ser intercambiadas
	posicion1 := 0
	posicion2 := 0
	//Recorremos el arreglo
	for i := posicion1; i < len(arreglo); i++ {
		min = arreglo[i] //Se asigna como mínimo el valor que se esta recorriendo
		//Se recorre el arreglo, esta vez desde la posición en la que se encuentra el mínimo actual
		for j := posicion2; j < len(arreglo); j++ {
			//En caso de que haya otro valor menor, se asigna como mínimo
			if min > arreglo[j] {
				min = arreglo[j]
				posicion2 = j
			}
		}

		//Se intercambian las posiciones, el mínimo con el valor que se había escogido antes como minimo
		arreglo[posicion2] = arreglo[posicion1]
		arreglo[posicion1] = min

		//Se aumenta la posicion y se asigna a la posicion2 el valor de la primera
		posicion1++
		posicion2 = posicion1

		//En caso de que se este en el ultimo el elemento se sale del ciclo
		if posicion1 == len(arreglo)-1 {
			break
		}
	}
	//Se retorna el arreglo
	return arreglo
}

// 5. Función que ordene ascendentemente un arreglo de enteros, mediante Quicksort
func ordenarQuicksort(arreglo []int) []int {

	//Se verifica si el arreglo esta vació, de ser asi se retorna un arreglo vació
	if len(arreglo) < 1 {
		return []int{}
	}

	//Se declaran los subarreglo que contendrán los numeros mayores y menores al pivote
	izq := []int{}
	der := []int{}
	pivote := arreglo[0] //Se usa el primer elemento como pivote

	//Con un ciclo for se verifican cuales numeros son menores y cuales mayores al pivote
	//Se guardan en sus respectivos arreglos
	for i := 1; i < len(arreglo); i++ {
		if arreglo[i] < pivote {
			izq = append(izq, arreglo[i])
		} else {
			der = append(der, arreglo[i])
		}
	}
	//Realiza llamada recursivas con los subarreglos que contienen los menores y mayores al pivote
	//De esta manera gracias a las llamadas recursivas se ira ordenando
	return append(append(ordenarQuicksort(izq), pivote), ordenarQuicksort(der)...)

}

// 6. Función que busca un valor entero en un arreglo no ordenado por medio de búsqueda secuencial
func busqueda_secuencial(arreglo []int, llave int) (bool, int) {
	//Variable donde se almacenan la cantidad de comparaciones hechas
	comparaciones := 0
	//Con un ciclo for se compara cada numero del arreglo con la llave
	//En caso de que sean iguales se retorna true y las comparaciones
	//En cada iteracion las comparaciones aumentan
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == llave {
			comparaciones++
			return true, comparaciones
		}
		comparaciones++
	}
	//Si no se encontro se retorna false y las comparaciones
	return false, comparaciones
}

// 7. Función que busca un valor entero en un arreglo ordenado por medio de busqueda binaria
func busqueda_binaria(arreglo []int, llave int) (bool, int) {

	//Asignamos valores para las posiciones que se necesitan comparar
	//En este caso se compara la del inicio, final, y una posicion media
	inicio := 0
	final := len(arreglo) - 1
	medio := 0
	//Variable para llevar registro de las comparaciones
	comparaciones := 0

	//Se realiza un ciclo for mientras la posicion del inicio sea menor a la del final
	for inicio <= final {
		//Se obtiene la posicion media
		medio = (inicio + final) / 2
		comparaciones++
		//Se realiza las comparaciones
		if arreglo[medio] == llave {
			return true, comparaciones
		} else if llave > arreglo[medio] {
			//En caso de que el numero a buscar sea mayor al valor que se encuentra en la media
			//La posicion del inicio ahora sera la media mas 1 para comparar despues del punto medio
			inicio = medio + 1
		} else if llave < arreglo[medio] {
			//En caso de que el numero a buscar sea menor al valor que se encuentra en la media
			//La posicion del final ahora sera la media menos 1 para comparar antes del punto medio
			final = medio - 1
		}
	}

	//En caso de salir del ciclo for significa que no encontro la llave
	//Por lo que retorna false y las comparaciones
	return false, comparaciones

}

// 8. Diseñar una representación para un Árbol Binario de Búsqueda
// Se usan las estructuras de Golang para definir la estructura del arbol de busqueda binario
type ArbolBB struct {
	numero  int      //Se define variable que guardara valor entero
	HijoIzq *ArbolBB //Se define un puntero a una estructura de tipo Arbol que sera nodo izquierdo
	HijoDer *ArbolBB //Se define un puntero a una estructura de tipo Arbol que sera nodo izquierdo
}

// 9. Funcion que realiza la insercion de un valor entero en un Arbol Binario de Busqueda
// Se retorna la cantidad de comparaciones realizadas hasta cuando se inserto o no
func InsertarNodo(Nodo **ArbolBB, llave int) int {
	comparaciones := 1 //Variable donde se almacenan las comparaciones
	//En caso de que el nodo este nulo, se inserta la llave en el arbol
	if *Nodo == nil {
		*Nodo = new(ArbolBB)
		(*Nodo).numero = llave
		(*Nodo).HijoDer = nil
		(*Nodo).HijoIzq = nil
	} else {
		//En caso de que la llave ya existe, no se inserta y se retorna las comparaciones
		if llave == (*Nodo).numero {
			return comparaciones
		} else {
			if llave > (*Nodo).numero {
				//Si la llave es mayor al la llave del nodo actual
				// Se vuelve a llamar la funcion ahora recorriendo el hijo derecho
				comparaciones += InsertarNodo(&(*Nodo).HijoDer, llave)
			} else {
				//Si la llave es mayor al la llave del nodo actual
				// Se vuelve a llamar la funcion ahora recorriendo el hijo derecho
				comparaciones += InsertarNodo(&(*Nodo).HijoIzq, llave)
			}
		}
	}

	return comparaciones //Se retorna las comparaciones
}

// 10. Funcion que realiza la busqueda de una llave en un Arbol Binario de Busqueda
// Se retrona un valor booleano (Si se encontro o no la llave) y la cantidad de comparaciones hechas
func BuscarNodo(Nodo *ArbolBB, llave int) (bool, int) {
	//Variables para almacenar las comparaciones y si se encontro o no
	comparaciones := 1
	comparacionesTemp := 0
	encontrado := false

	//En caso de que el nodo este nulo significa que se recorrio y no se encontro
	if Nodo == nil {
		return encontrado, comparaciones
	} else {

		comparacionesTemp++ //Las comparaciones temporales aumentan en 1

		//Si la llave es igual a la llave del nodo actual significa que si se encontro
		if llave == Nodo.numero {
			encontrado = true
			return encontrado, comparaciones
		} else {
			if llave > Nodo.numero {
				//Si la llave es mayor al la llave del nodo actual
				// Se vuelve a llamar la funcion ahora recorriendo el hijo derecho
				// A las comparaciones se le suma las comparaciones temporales que tendran guardadas las comparaciones totales
				encontrado, comparacionesTemp = BuscarNodo(Nodo.HijoDer, llave)
				comparaciones += comparacionesTemp

			} else {
				//Si la llave es mayor al la llave del nodo actual
				// Se vuelve a llamar la funcion ahora recorriendo el hijo izquierdo
				// A las comparaciones se le suma las comparaciones temporales que tendran guardadas las comparaciones totales
				encontrado, comparacionesTemp = BuscarNodo(Nodo.HijoIzq, llave)
				comparaciones += comparacionesTemp
			}
		}
	}
	return encontrado, comparaciones //Se retorna las comparaciones y el booleano que indica si se encontro o no
}

// Funciones complementarias
// Altura de ABB
func AlturaArbol(Nodo *ArbolBB) float64 {
	//Si esta vacio se retorna 0
	if Nodo == nil {
		return 0
	} else {
		//Se le suma uno al maximo de entre la altura del hijo izquierdo y el derecho
		return 1 + math.Max(AlturaArbol(Nodo.HijoIzq), AlturaArbol(Nodo.HijoDer))
	}
}

// Tamaño de Abb
func sizeArbol(Nodo *ArbolBB) float64 {
	//Si esta vacio se retorna 0
	if Nodo == nil {
		return 0
	} else {
		//Se le suma uno al tamaño del hijo izquierdo y el hijo derecho
		return 1 + sizeArbol(Nodo.HijoIzq) + sizeArbol(Nodo.HijoDer)
	}
}

// Densidad de ABB
func DensidadArbol(Nodo *ArbolBB) float64 {
	//Se divida el tamaño entre la altura
	return sizeArbol(Nodo) / AlturaArbol(Nodo)
}

//#########################################EXPERIMENTOS############################################
//En el main se realizan 4 experimentos
// Se usaran 4 valores distintos para n, estos son: 200,1000,3000,5000

func main() {

	//Variables que se usaras a lo largo de los experimentos
	var n int
	var arreglo []int
	var A []int
	var TS []int
	var TOS []int
	var TOQ []int
	var Abb *ArbolBB
	var comparaciones int
	var comparacionesTemp int
	var encontrado bool
	fmt.Println(encontrado) //Ignorar

	// EXPERIMENTO 1
	//  n = 200
	fmt.Println("EXPERIMENTO 1")

	//a. Crear arreglo A de tamaño n
	fmt.Println("Inicio Punto A: Generar arreglo de numeros pseudoaleatorios de tamaño n=200")
	n = 200
	A = generarPseudoAleatorios(n, 45) //Se usa s0=45 puede ser cualquier valor entre 11 y 257
	fmt.Println("Arreglo A generado:")
	fmt.Println(A)
	fmt.Println("Fin Punto A")

	
	//Falta b y c
	//b.
	//c.


	//d. Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion
	fmt.Println("Inicio Punto D: Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion")
	TS = []int{}
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += insertarNumeros(&TS, A[i])
	}
	fmt.Println("Arreglo TS generado:")
	fmt.Println(TS)
	fmt.Println("Comparaciones totales hechas para las inserciones en TS:",comparaciones)
	fmt.Println("Fin Punto D")



	//e. Crear arreglo TOS que es una copia de A y ordenarla por seleccion
	fmt.Println("Inicio Punto E: En un arreglo TOS copiar elementos de A y ordenarlos por seleccion")
	TOS = make([]int, len(A))
	copy(TOS, A)
	TOS = ordenarSeleccion(TOS)
	fmt.Println("Arreglo TOS Ordenado:")
	fmt.Println(TOS)
	fmt.Println("Fin Punto E")



	//f. Crear arreglo TOQ que es una copia de A y ordenarla por quicksort
	fmt.Println("Inicio Punto F: En un arreglo TOQ copiar elementos de A y ordenarlos por quicksort")
	TOQ = make([]int, len(A))
	copy(TOQ, A)
	TOQ = ordenarQuicksort(TOQ)
	fmt.Println("Arreglo TOQ Ordenado:")
	fmt.Println(TOQ)
	fmt.Println("Fin Punto F")



	//g. Crear arbol Abb, e insertar los elementos de A
	fmt.Println("Inicio Punto G: Insertar en un arbol Abb los elementos de A")
	Abb = new(ArbolBB)
	Abb = nil
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += InsertarNodo(&Abb, A[i])
	}
	fmt.Println("Comparaciones totales hechas para las inserciones en Abb:",comparaciones)
	fmt.Println("Fin Punto G")



	//i. Generar un arreglo con 10000 numeros aleatorios y guardar las estadisticas
	fmt.Println("Inicio Punto I: Crear arreglo con 10000 numeros y buscar los numeros en TS,TOS,TOQ,Abb")
	arreglo = generarPseudoAleatorios(10000, 45)

	//Buscar numeros del arreglo en TS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_secuencial(TS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TS:",comparaciones)


	//Buscar numeros del arreglo en TOS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOS:",comparaciones)


	//Buscar numeros del arreglo en TOQ
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOQ:",comparaciones)


	//Buscar numeros del arreglo en Abb
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = BuscarNodo(Abb, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en Abb:",comparaciones)



	//Altura de Abb
	fmt.Println("Altura de Abb")
	fmt.Println(AlturaArbol(Abb))

	//Densidad de Abb
	fmt.Println("Densidad de Abb")
	fmt.Println(DensidadArbol(Abb))










	// EXPERIMENTO 2
	//  n = 1000
	fmt.Println("EXPERIMENTO 2")

	//a. Crear arreglo A de tamaño n
	fmt.Println("Inicio Punto A: Generar arreglo de numeros pseudoaleatorios de tamaño n=1000")
	n = 1000
	A = generarPseudoAleatorios(n, 45) //Se usa s0=45 puede ser cualquier valor entre 11 y 257
	fmt.Println("Arreglo A generado:")
	fmt.Println(A)
	fmt.Println("Fin Punto A")

	
	//Falta b y c
	//b.
	//c.


	//d. Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion
	fmt.Println("Inicio Punto D: Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion")
	TS = []int{}
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += insertarNumeros(&TS, A[i])
	}
	fmt.Println("Arreglo TS generado:")
	fmt.Println(TS)
	fmt.Println("Comparaciones totales hechas para las inserciones en TS:",comparaciones)
	fmt.Println("Fin Punto D")



	//e. Crear arreglo TOS que es una copia de A y ordenarla por seleccion
	fmt.Println("Inicio Punto E: En un arreglo TOS copiar elementos de A y ordenarlos por seleccion")
	TOS = make([]int, len(A))
	copy(TOS, A)
	TOS = ordenarSeleccion(TOS)
	fmt.Println("Arreglo TOS Ordenado:")
	fmt.Println(TOS)
	fmt.Println("Fin Punto E")



	//f. Crear arreglo TOQ que es una copia de A y ordenarla por quicksort
	fmt.Println("Inicio Punto F: En un arreglo TOQ copiar elementos de A y ordenarlos por quicksort")
	TOQ = make([]int, len(A))
	copy(TOQ, A)
	TOQ = ordenarQuicksort(TOQ)
	fmt.Println("Arreglo TOQ Ordenado:")
	fmt.Println(TOQ)
	fmt.Println("Fin Punto F")



	//g. Crear arbol Abb, e insertar los elementos de A
	fmt.Println("Inicio Punto G: Insertar en un arbol Abb los elementos de A")
	Abb = new(ArbolBB)
	Abb = nil
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += InsertarNodo(&Abb, A[i])
	}
	fmt.Println("Comparaciones totales hechas para las inserciones en Abb:",comparaciones)
	fmt.Println("Fin Punto G")



	//i. Generar un arreglo con 10000 numeros aleatorios y guardar las estadisticas
	fmt.Println("Inicio Punto I: Crear arreglo con 10000 numeros y buscar los numeros en TS,TOS,TOQ,Abb")
	arreglo = generarPseudoAleatorios(10000, 45)

	//Buscar numeros del arreglo en TS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_secuencial(TS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TS:",comparaciones)


	//Buscar numeros del arreglo en TOS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOS:",comparaciones)


	//Buscar numeros del arreglo en TOQ
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOQ:",comparaciones)


	//Buscar numeros del arreglo en Abb
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = BuscarNodo(Abb, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en Abb:",comparaciones)
	fmt.Println("Fin Punto I")



	//Altura de Abb
	fmt.Println("Altura de Abb")
	fmt.Println(AlturaArbol(Abb))

	//Densidad de Abb
	fmt.Println("Densidad de Abb")
	fmt.Println(DensidadArbol(Abb))










	// EXPERIMENTO 3
	//  n = 3000
	fmt.Println("EXPERIMENTO 3")

	//a. Crear arreglo A de tamaño n
	fmt.Println("Inicio Punto A: Generar arreglo de numeros pseudoaleatorios de tamaño n=3000")
	n = 3000
	A = generarPseudoAleatorios(n, 45) //Se usa s0=45 puede ser cualquier valor entre 11 y 257
	fmt.Println("Arreglo A generado:")
	fmt.Println(A)
	fmt.Println("Fin Punto A")

	
	//Falta b y c
	//b.
	//c.


	//d. Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion
	fmt.Println("Inicio Punto D: Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion")
	TS = []int{}
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += insertarNumeros(&TS, A[i])
	}
	fmt.Println("Arreglo TS generado:")
	fmt.Println(TS)
	fmt.Println("Comparaciones totales hechas para las inserciones en TS:",comparaciones)
	fmt.Println("Fin Punto D")



	//e. Crear arreglo TOS que es una copia de A y ordenarla por seleccion
	fmt.Println("Inicio Punto E: En un arreglo TOS copiar elementos de A y ordenarlos por seleccion")
	TOS = make([]int, len(A))
	copy(TOS, A)
	TOS = ordenarSeleccion(TOS)
	fmt.Println("Arreglo TOS Ordenado:")
	fmt.Println(TOS)
	fmt.Println("Fin Punto E")



	//f. Crear arreglo TOQ que es una copia de A y ordenarla por quicksort
	fmt.Println("Inicio Punto F: En un arreglo TOQ copiar elementos de A y ordenarlos por quicksort")
	TOQ = make([]int, len(A))
	copy(TOQ, A)
	TOQ = ordenarQuicksort(TOQ)
	fmt.Println("Arreglo TOQ Ordenado:")
	fmt.Println(TOQ)
	fmt.Println("Fin Punto F")



	//g. Crear arbol Abb, e insertar los elementos de A
	fmt.Println("Inicio Punto G: Insertar en un arbol Abb los elementos de A")
	Abb = new(ArbolBB)
	Abb = nil
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += InsertarNodo(&Abb, A[i])
	}
	fmt.Println("Comparaciones totales hechas para las inserciones en Abb:",comparaciones)
	fmt.Println("Fin Punto G")



	//i. Generar un arreglo con 10000 numeros aleatorios y guardar las estadisticas
	fmt.Println("Inicio Punto I: Crear arreglo con 10000 numeros y buscar los numeros en TS,TOS,TOQ,Abb")
	arreglo = generarPseudoAleatorios(10000, 45)

	//Buscar numeros del arreglo en TS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_secuencial(TS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TS:",comparaciones)


	//Buscar numeros del arreglo en TOS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOS:",comparaciones)


	//Buscar numeros del arreglo en TOQ
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOQ:",comparaciones)


	//Buscar numeros del arreglo en Abb
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = BuscarNodo(Abb, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en Abb:",comparaciones)
	fmt.Println("Fin Punto I")



	//Altura de Abb
	fmt.Println("Altura de Abb")
	fmt.Println(AlturaArbol(Abb))

	//Densidad de Abb
	fmt.Println("Densidad de Abb")
	fmt.Println(DensidadArbol(Abb))










	// EXPERIMENTO 4
	//  n = 5000
	fmt.Println("EXPERIMENTO 4")

	//a. Crear arreglo A de tamaño n
	fmt.Println("Inicio Punto A: Generar arreglo de numeros pseudoaleatorios de tamaño n=5000")
	n = 5000
	A = generarPseudoAleatorios(n, 45) //Se usa s0=45 puede ser cualquier valor entre 11 y 257
	fmt.Println("Arreglo A generado:")
	fmt.Println(A)
	fmt.Println("Fin Punto A")

	
	//Falta b y c
	//b.
	//c.


	//d. Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion
	fmt.Println("Inicio Punto D: Insertar en un arreglo TS los elementos de A mediante algoritmo de insercion")
	TS = []int{}
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += insertarNumeros(&TS, A[i])
	}
	fmt.Println("Arreglo TS generado:")
	fmt.Println(TS)
	fmt.Println("Comparaciones totales hechas para las inserciones en TS:",comparaciones)
	fmt.Println("Fin Punto D")



	//e. Crear arreglo TOS que es una copia de A y ordenarla por seleccion
	fmt.Println("Inicio Punto E: En un arreglo TOS copiar elementos de A y ordenarlos por seleccion")
	TOS = make([]int, len(A))
	copy(TOS, A)
	TOS = ordenarSeleccion(TOS)
	fmt.Println("Arreglo TOS Ordenado:")
	fmt.Println(TOS)
	fmt.Println("Fin Punto E")



	//f. Crear arreglo TOQ que es una copia de A y ordenarla por quicksort
	fmt.Println("Inicio Punto F: En un arreglo TOQ copiar elementos de A y ordenarlos por quicksort")
	TOQ = make([]int, len(A))
	copy(TOQ, A)
	TOQ = ordenarQuicksort(TOQ)
	fmt.Println("Arreglo TOQ Ordenado:")
	fmt.Println(TOQ)
	fmt.Println("Fin Punto F")



	//g. Crear arbol Abb, e insertar los elementos de A
	fmt.Println("Inicio Punto G: Insertar en un arbol Abb los elementos de A")
	Abb = new(ArbolBB)
	Abb = nil
	comparaciones = 0
	for i := 0; i < len(A); i++ {
		comparaciones += InsertarNodo(&Abb, A[i])
	}
	fmt.Println("Comparaciones totales hechas para las inserciones en Abb:",comparaciones)
	fmt.Println("Fin Punto G")



	//i. Generar un arreglo con 10000 numeros aleatorios y guardar las estadisticas
	fmt.Println("Inicio Punto I: Crear arreglo con 10000 numeros y buscar los numeros en TS,TOS,TOQ,Abb")
	arreglo = generarPseudoAleatorios(10000, 45)

	//Buscar numeros del arreglo en TS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_secuencial(TS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TS:",comparaciones)


	//Buscar numeros del arreglo en TOS
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOS:",comparaciones)


	//Buscar numeros del arreglo en TOQ
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = busqueda_binaria(TOS, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en TOQ:",comparaciones)


	//Buscar numeros del arreglo en Abb
	comparaciones = 0
	comparacionesTemp = 0
	for i := 0; i < len(arreglo); i++ {
		encontrado, comparacionesTemp = BuscarNodo(Abb, arreglo[i])
		comparaciones += comparacionesTemp
	}
	fmt.Println("Comparaciones totales hechas para la busqueda de valores en Abb:",comparaciones)
	fmt.Println("Fin Punto I")



	//Altura de Abb
	fmt.Println("Altura de Abb")
	fmt.Println(AlturaArbol(Abb))

	//Densidad de Abb
	fmt.Println("Densidad de Abb")
	fmt.Println(DensidadArbol(Abb))


}








	