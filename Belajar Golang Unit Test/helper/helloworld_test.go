package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
//Table Benchmark
func BenchmarkTable(b*testing.B){
	benchmarks:=[]struct{
		name string
		request string
	}{
		{
			name: "Eko",
			request: "Eko",
		},
		{
			name: "Jason",
			request: "Jason",
		},
	}
	for _, benchmark := range benchmarks{
		b.Run(benchmark.name,func(b *testing.B) {
			for b.Loop(){
				HelloWorld(benchmark.request)
			}
		})
	}
}

//SubBenchmark

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for b.Loop() {
			HelloWorld("Eko")
		}
	})
	b.Run("Jason", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jason")
		}
	})
}

//BenchMark

func BenchmarkHelloWorldEko(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldJason(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Jason")
	}
}

// Table Test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorldEko",
			request:  "Eko",
			expected: "Hello World Eko",
		},
		{
			name:     "HelloWorldJason",
			request:  "Jason",
			expected: "Hello World Jason",
		},
		{
			name:     "HelloWorldDennis",
			request:  "Dennis",
			expected: "Hello World Dennis",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, result, test.expected)
		})
	}
}

// Function SubTest
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello World Eko", result, "Result must be 'Hello World Eko'")
	})
	t.Run("Jason", func(t *testing.T) {
		result := HelloWorld("Jason")
		require.Equal(t, "Hello World Jason", result, "Result must be 'Hello World Jason'")
	})
}

// before and after
func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")
	m.Run()
	//after
	fmt.Println("After Unit Test")
}

// Skip this test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Window OS")
	}
	result := HelloWorld("Eko")
	require.Equal(t, "Hello World Eko", result, "Result must be 'Hello World Eko'")
}

//Assertion and Require from testify

// require use FailNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello World Eko", result, "Result must be 'Hello World Eko'")
	fmt.Println("Test HelloWorld with Require done")
}

// assert use Fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello World Eko", result, "Result must be 'Hello World Eko'")
	fmt.Println("Test HelloWorld with Assert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello World Eko" {
		//error
		t.Error(" Expected 'Hello World Eko' but got ", result)
	}
	fmt.Println("Test HelloWorld done")
}

func TestHelloWorldJason(t *testing.T) {
	result := HelloWorld("Jason")
	if result != "Hello World Jason" {
		//error
		t.Fatal(" Expected 'Hello World Jason' but got ", result)
	}
	fmt.Println("Test HelloWorldJason done")
}

// got test (run unit test command)

// go test -v (run all unit test with verbose)

// go test -v -run=[unit test function name] (run specific unit test function)

// go test  -run=[unit test function name]/[unit test sub test name] (run specific sub test function from unit test function)

// go test -v ./... (run all directory that have unit test)

// go test -v -bench=.(run all unit test and benchmark)

// go test -v -run=[unit test function/ benchmark name] -bench=. (run spesific name unit test and all benchmark)
// go test -v -run=[not unit test function name] -bench=. (run  all benchmark only)
// got test -v -run[not unit test function name] -bench=[benchmark function name] (run only speficic name of benchmark)
// got test -v -bench=. ./... (run all benchmarck and unit test from root)
// got test -v -run[not unit test function name] -bench=. ./... (run all benchmarck test from root)
