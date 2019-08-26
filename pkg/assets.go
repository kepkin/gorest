// +build dev

package pkg

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("../assets")
