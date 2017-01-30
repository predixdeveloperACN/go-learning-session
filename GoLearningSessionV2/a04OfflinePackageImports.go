
// NOTE: ALWAYS place your Packages inside your Go Root's "src" folder -- NEVER the "pkg" folder!
//       The pkg folder is reserved for Go Packages that you install from online sources, such as Github.
//       It's also where the Go compiler will sometimes place copies of your packages to avoid recompiling.
//
// -------------------------------------------------------

// The following code is located in:
//
// D:\VinceGo\src\samplepkg1\file1.go
//
//     package samplepkg1
//
//     import (
//         "fmt"
//     )
//
//     func Hello() {
//         fmt.Println("Hello")
//     }
//
// -------------------------------------------------------

// The following code is located in:
//
// D:\VinceGo\src\samplepkg2\subpkg\file2.go
//
//     package subpkg
//
//     import (
//         "fmt"
//     )
//
//     func World() {
//         fmt.Println("World")
//     }
//
// -------------------------------------------------------

// This is our main Go program, which can be located anywhere within your Go Root's "src" folder.
// For this example, our .go file is located in:
//
// D:\VinceGo\src\GoLearningSession

package main

// Note the aliases we give our packages during import.
// We can use them instead of the actual package name during coding.
//
import (
    p1 "samplepkg1"
    p2 "samplepkg2/subpkg"
)

func main(){

    p1.Hello()
    p2.World()
    // OUTPUT:
    // Hello
    // World

    // NOTE 1: If we didn't use aliases one our imports,
    //         we would need to use the original Package Name instead:
    //
    //     samplepkg1.Hello()
    //     subpkg.World()
    //
    // NOTE 2: You can't use the original Package Name if you defined an Alias for it.
    //
    // NOTE 3: Go has no concept of Subpackages. Subpackages are just
    //         directories inside another directory -- which is why
    //         you MUST Import Subpackages explicitly. You won't be able
    //         to use Subpackages by importing only the parent package:
    //
    //         e.g. Importing only "math" will not let you use "math.rand" or "math/rand".
    //              You MUST explicitly import "math/rand" as well.
    //
    //         This is because "rand" is not a child of the "math" package,
    //         even if it's located inside it. There is NO relationship between the two.
}
