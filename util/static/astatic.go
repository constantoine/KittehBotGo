// Code generated by fileb0x at "2017-10-16 12:07:55.933157866 +0100 BST m=+0.044623998" from config file "b0x.yaml" DO NOT EDIT.

package static


import (
  "bytes"
  "compress/gzip"
  "io"
  "log"
  "net/http"
  "os"
  "path"

  "golang.org/x/net/webdav"
  "golang.org/x/net/context"


)

var ( 
  // CTX is a context for webdav vfs
  CTX = context.Background()

  
  // FS is a virtual memory file system
  FS = webdav.NewMemFS()
  

  // Handler is used to server files through a http handler
  Handler *webdav.Handler

  // HTTP is the http file system
  HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {}



// FileAssetsGraphsJs is "assets/graphs.js"
var FileAssetsGraphsJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x56\x5b\x6f\xe3\x36\x13\x7d\xcf\xaf\x98\x8f\xbb\x80\xe5\xac\x20\xcb\xc9\x3a\x17\x05\xfe\x8a\x26\x9b\x2c\x5a\x34\xc0\xa2\x71\x8a\x16\x69\x1e\x68\x69\xec\xb0\xa1\x48\x83\xa4\x12\xbb\x81\xfe\x7b\x41\x8a\xba\xf8\x86\x5d\x2c\xfa\x50\xf8\xc1\xe0\xcc\x1c\xf2\x90\x3c\x47\x9c\xf7\xc1\xac\x10\xa9\x61\x52\x40\xd0\x87\xb7\x03\x80\x57\x26\x32\xf9\x1a\x65\x0b\x0d\x63\x78\x78\xbc\x68\x43\xcb\xdf\x28\x87\x31\x0c\x3b\xa1\x55\x15\x8a\x3b\xa1\x62\x91\x51\x83\x3f\x09\x83\xea\xa5\x4e\x1e\x00\xbc\x8f\xe8\x5f\x74\x19\xd8\x15\x00\xcc\x6a\x81\x09\xf4\x3e\x5f\x4f\x7a\xa1\x0b\x14\x8a\x27\xd0\x1b\x30\x8f\xf2\xd1\x54\xe6\x0b\x8e\x06\x13\x68\x49\x2a\xd4\x0b\x29\x34\x4e\x70\x69\x2a\xc2\x55\xa5\xd0\x92\x63\xc4\xe5\x7c\xbd\xc2\xe7\xf7\x71\x5b\x50\xa5\xf1\x86\x4b\x6a\xd6\x60\xd1\xf6\x1c\xe5\x01\x40\xd9\x77\x5b\xf1\x93\xe5\x98\xfb\x43\x7a\x83\x65\x02\x71\x08\xab\x04\x62\x28\xdd\x91\xbd\x50\x05\x5a\xa6\xcf\x68\x60\x0c\x4c\x06\x15\xf2\x8a\x8a\x17\xaa\x7f\xbe\x8b\x68\x96\x5d\x49\x2e\xd5\x1d\x9a\x80\xa4\x85\x36\x32\xaf\xc7\x24\x84\x07\xb7\x22\x79\x77\x7c\x7e\x3c\x3b\x39\x26\xa1\x1f\xe2\x28\x3b\xbb\x8c\x9b\xe1\x6c\x36\x3d\x3e\x39\x6d\x87\xe7\x67\x1f\x4f\x86\xcd\x30\x3b\x3f\x39\x1f\xcd\x5a\x6c\x3c\x3a\x1b\xc5\xe4\x00\xe0\x71\x73\x13\xf7\x9a\xce\x11\xc6\x20\xf0\xb5\x25\x78\xf5\x44\x95\x09\x48\x9d\x26\xa1\x3f\x6a\x2a\x58\x4e\xed\x4d\x5c\x0b\x3a\xe5\x98\x25\x60\x54\x81\xe1\x7a\xae\x1b\xf4\x27\xc9\x5e\xb0\x1b\x9d\xd2\xf4\x79\xae\x64\x21\xaa\x63\x48\x80\x18\x45\x85\x5e\x50\x85\xc2\x78\xd2\x74\xc9\xf4\xef\x49\x73\xc5\x9c\x4e\x91\xdf\x48\x61\x6a\xc4\xbb\xd3\xa1\xfd\xf9\x6a\x00\xce\x04\x36\x39\x7a\x64\x7f\x4d\xce\xb0\xf4\x79\xff\x4a\x36\x6f\x38\x26\x40\x26\x2c\x47\x08\x34\xa6\x52\x64\xfa\x87\x3e\xf1\xe9\x4a\x01\x2d\xad\x3f\x5a\x5a\x73\xc5\xb2\xc9\x13\x4b\x9f\x05\x6a\x6d\x55\xf0\x6f\xd3\xdd\xcc\xe5\x4c\xb0\xbc\xc8\x13\x38\xda\x24\x7f\x8b\xb9\x54\x2b\xb8\xd7\x98\x41\x70\x7b\xe9\xd9\x7b\xda\x46\x4a\x3e\x61\x8b\x04\xde\x20\x95\x4a\xa0\xfa\x95\x66\xac\x70\x8c\x61\x26\x85\xb9\x33\x2b\x37\x89\x90\x2a\xa7\x9c\xd4\xb0\x8c\x1a\x9a\x58\x85\xa7\x35\x9d\xab\x4f\xa3\xd3\x8f\x31\x09\xbd\x89\x89\xdd\x07\x09\x5d\xe1\x17\xc9\x84\xd1\x09\x58\x5b\x94\x8f\xde\x30\x5b\x52\x8b\x14\x8a\x0c\x55\xb0\x91\xbb\x33\xd4\xe8\xfd\x32\x74\xe9\x6d\x19\x7e\x2a\x94\x97\xdc\x4e\x8b\x87\x5f\x15\xed\xb7\x28\x31\xf5\xd6\x4c\x60\xcb\xab\xdf\xa7\xd4\xba\xe0\x8e\xfd\x8d\x09\x0c\xcf\xba\x9a\xd8\x25\x26\xab\x87\x9d\x22\xfb\x0e\xc9\xff\x27\xb5\x1c\xef\xd1\xf2\x96\x8c\xbd\x1e\x7d\x75\x4d\x1c\x80\x89\x0c\x97\xbf\x7c\x8d\xe7\x66\xdd\x0d\xcd\x19\x5f\xd9\x6b\xa5\x9c\x4d\x15\xdb\x5b\xb8\x79\x51\xdd\xfc\x17\x4e\x53\xcc\x51\x58\x79\xc8\xc2\x68\x96\xe1\x9e\x79\x54\x4e\x8d\x41\xd5\x7d\xce\xb0\xdf\xd9\x84\xfd\x5a\x9a\x42\x09\xc0\xa8\xb1\x53\xb4\x8a\x8c\xbc\x61\x4b\xcc\x82\x51\x1f\x3e\x00\x81\xdb\x4b\x72\xd1\x40\xca\x76\x25\xef\xc7\x29\x55\x9d\xe5\xbb\xb6\x5c\x7b\xb7\x7c\x45\xe9\xfe\x77\x79\xd5\x19\x6e\xcd\xab\x83\x43\xff\xb0\x55\x1e\x73\xde\x84\x71\xb3\x99\xa0\xde\xca\xda\x53\xef\xdf\x76\x32\x47\x63\xb9\xb4\xb2\xa8\xd8\x7e\xbe\x9e\x34\x21\x5d\xa4\xa9\x13\x5e\x33\xa3\x45\xf4\xe1\x2d\xe8\x1f\x0e\x5a\x6e\x1a\x8d\xfd\x4c\xcb\xc2\x6c\xf4\x2e\x6e\x9a\xea\xd1\x8d\xa4\x08\x7a\x39\xe6\xa0\xed\x2e\x7a\x61\xe7\xc4\x73\x3d\xdf\xdd\x37\x10\x6d\x8a\xd9\xec\x7f\xa4\xee\x19\xdc\x13\x6e\x43\x30\x86\x5c\xcf\x23\xbd\xe0\xcc\x04\xe4\x4f\x41\xfa\x17\xeb\x6d\x85\xef\x82\x3a\xcd\x84\xc3\x3d\xc4\x8f\x4d\x65\xb6\xd0\xd1\xa2\xd0\x4f\x81\xeb\x15\x6c\x27\xe5\xda\x05\x87\x2c\x9b\x2a\x36\x83\xa0\xdb\x6c\xfd\x1f\x8e\xe2\xae\x40\xda\xde\x2c\xd2\x4f\x6c\x66\x82\x06\x59\xae\x33\xaa\xee\xd8\x51\x4e\x31\x88\x43\xf0\x01\x8e\x62\x6e\x9e\x36\xf9\xfb\xa4\xe3\xd7\xac\xf5\x66\xf9\xed\xda\x52\x58\x7d\x0f\x12\x20\xf7\x9a\x89\x39\xe9\x6a\x70\x37\x68\xd8\x05\xfd\xc8\xb9\x4c\xa9\xc1\xec\x1b\x80\x47\x5d\xe0\x15\x47\x2a\x1c\xcc\xa3\x76\x6c\x63\xfb\x81\x59\xcf\x6f\x8a\x7a\xbd\xc0\x9d\xf9\x87\xf1\x9e\x6e\x71\x00\xc3\x38\x8e\xab\x39\xab\x2b\x2b\x43\x17\xaa\xcc\x31\x28\x9d\x45\x6c\xa6\xbc\x70\x7a\x75\x7e\xd1\x68\xea\x19\x82\x2d\xa7\x78\xb1\x62\xce\x4c\x25\xd7\x39\x9a\x5e\xe8\xcc\x42\xea\x15\x76\x92\x71\xc9\xc3\xc1\x81\x5d\xed\x9f\x00\x00\x00\xff\xff\xc2\xe6\x1e\xdc\xc3\x0b\x00\x00")

// FileAssetsStyleCSS is "assets/style.css"
var FileAssetsStyleCSS = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x53\xdd\x8e\x9b\x3c\x10\x7d\x15\x3e\x45\x9f\xd4\x4a\x31\x32\xbb\x09\xed\x9a\x9b\xf6\xae\xf7\x55\x1f\xc0\xe0\x01\xa6\x6b\x3c\x96\x3d\x64\xd9\x22\xde\xbd\x22\x81\x2c\xd9\x9f\x07\x68\x2e\xa2\x31\x3e\x73\xe6\xf8\xf8\xb8\x24\xf3\x3c\x96\xba\x7a\x6c\x02\xf5\xce\x88\x8a\x2c\x05\xb5\xab\xeb\xba\xa8\xc9\xb1\xa8\x75\x87\xf6\x59\x05\x2a\x89\x69\xdf\x82\x3d\x01\x63\xa5\x13\x07\x3d\xec\x7f\xac\xcb\xfd\xf7\x80\xda\xee\xa3\x76\x51\x44\x08\xb8\x34\x47\xfc\x03\x2a\x3b\xf8\xa1\xf0\xda\x18\x74\x8d\x60\xf2\x2a\x3f\xf8\x61\xda\xb5\xa0\x0d\x84\x77\x46\xdf\x1f\x0e\x0f\x47\x28\x5a\xc0\xa6\xe5\x33\xb8\xe8\x74\x68\xd0\x89\x92\x98\xa9\x53\x72\x65\x9b\x2b\x8a\xc8\x48\x4e\xd5\x38\x80\x59\x59\x93\xb4\x22\xc7\x1a\x1d\x04\x51\xdb\x1e\xcd\xb8\x76\x7c\x7d\x19\x9d\xa4\x4e\x9f\x4a\x1d\xc4\xa2\xe4\x32\x44\xc9\xeb\xfe\x2e\xa2\x81\x19\xc0\xd4\x34\x16\x44\xd9\x33\x93\x1b\x37\x0e\x55\x7d\x88\x14\x94\x27\x74\x0c\xa1\x30\x18\xbd\xd5\xcf\x0a\x9d\x45\x07\xa2\xb4\x54\x3d\x6e\x8d\xf8\x72\x3d\x8a\x92\x49\x76\x7c\xb1\x45\x65\x77\x7e\x48\xf2\xad\xb6\x32\x68\x67\xc6\x77\x19\x6f\x9a\xb2\x7c\xfe\x9b\x2b\xf9\xaa\x39\xd1\xe3\xeb\xcb\x3c\xcb\xb8\x9b\x1d\x3d\xd3\x2d\x16\x67\x05\xc3\xc0\xc2\x40\x45\x41\x9f\xcd\x74\xe4\x60\x4a\x6f\x8f\x7f\xd5\x32\x6f\xfe\x87\x9d\xa7\xc0\xda\xf1\x94\x7a\xdd\x80\x98\xfd\x06\xc7\x2b\x76\xb9\x30\x0b\x35\x2b\x74\x2d\x04\xe4\x4d\xcb\xb7\x0e\x0c\xea\x4f\x1d\x3a\xf1\x84\x86\x5b\x95\xe5\x52\xfa\xe1\xf3\xb8\x3a\x7e\x33\xeb\x43\x21\x67\x37\x36\xb4\xd3\xae\x83\xee\x27\x6b\x8e\xfb\xb9\xfa\x15\x75\x03\xe3\x72\xc6\xbb\xa3\xf4\xc3\x94\x56\x3a\x98\xb7\x91\x5b\x14\x16\x25\x05\x03\x41\xc9\x0b\x4e\x30\xb2\x85\xf1\xf2\x71\xcd\x5e\xe6\x87\x24\x92\x45\x73\x4d\xf4\xba\x21\xfd\xb0\x35\x39\x5f\x97\x4f\x17\x01\xf7\x52\x4e\xbb\x00\x27\x70\x3d\x08\xd6\xcd\x3e\xad\xfa\xc8\xd4\xcd\xf5\x3b\x8f\xc0\x3c\xe4\x0f\xc7\x7a\x51\x24\x82\x36\xd8\x47\x25\x6f\x18\x0f\x52\x16\x27\x08\xf3\xf3\xb3\x42\x5b\x6c\x9c\x62\xf2\xd3\x96\x77\x93\x3d\x29\xff\x9f\x76\x27\x8c\xc8\x14\xa2\xa8\x5a\x1d\x58\xcc\x63\xd7\x54\xaf\xa1\x4a\xef\x03\x74\x49\x9a\x07\xe8\xde\xe0\x67\xca\xeb\x83\xd3\x65\x24\xdb\x33\xcc\x66\xb9\x93\x8e\xbf\x57\x54\x15\xc0\x20\x7f\x14\x97\xe4\xdf\xfd\xfd\x0d\x00\x00\xff\xff\x68\x17\x6a\x3f\x2b\x05\x00\x00")

// FileIndexHTML is "index.html"
var FileIndexHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xbc\x56\x4d\x8f\xe3\x36\x0f\x3e\x3b\x40\xfe\x83\x46\xef\x75\x25\x65\x66\xf2\xa2\xc5\xae\x9d\xa2\xd8\xdd\x43\x0f\x45\x8b\xb6\x0b\xb4\x47\xda\x62\x6c\x65\x64\xc9\x2b\x2a\x5f\x28\xfa\xdf\x0b\xd9\x4e\xe2\xa4\x33\xc0\x4e\x5b\x34\x87\x58\x14\x1f\x3e\xfc\x08\x49\x27\xbf\xfb\xf0\xc3\xfb\x5f\x7e\xfb\xf1\x23\x6b\x62\x6b\x57\xf3\x59\x9e\x9e\xcc\x82\xab\x0b\x8e\x8e\xaf\xe6\xb3\x74\x87\xa0\x57\xf3\x59\x96\xb7\x18\x81\x55\x0d\x04\xc2\x58\xf0\x6d\x5c\x8b\xaf\xf9\x45\xd1\xc4\xd8\x09\xfc\xbc\x35\xbb\x82\xff\x2a\x3e\x7d\x2b\xde\xfb\xb6\x83\x68\x4a\x8b\x9c\x55\xde\x45\x74\xb1\xe0\xdf\x7d\x2c\x50\xd7\x38\xb1\x73\xd0\x62\xc1\x77\x06\xf7\x9d\x0f\x71\x02\xdd\x1b\x1d\x9b\x42\xe3\xce\x54\x28\x7a\xe1\x0d\x33\xce\x44\x03\x56\x50\x05\x16\x8b\xfb\x21\xc2\x2c\x8f\x26\x5a\x5c\x7d\x00\x6a\x4a\x0f\x41\xe7\x6a\xb8\x48\x4a\xc6\x72\x6b\xdc\x13\x6b\x02\xae\x0b\x9e\x82\xa4\xb7\x4a\xad\xbd\x8b\x24\x6b\xef\x6b\x8b\xd0\x19\x92\x95\x6f\x55\x45\xf4\xcd\x1a\x5a\x63\x8f\xc5\x4f\xbe\xf4\xd1\x73\x16\xd0\x16\x9c\xe2\xd1\x22\x35\x88\x71\x08\xbb\xe7\xbb\xd5\xdc\x38\x68\xe1\x50\x69\x27\x4b\xef\x23\xc5\x00\x5d\x12\x92\x8f\xf3\x85\x5a\xca\x85\x5c\x08\xb0\x5d\x03\xf2\x31\xf9\xbe\xe8\x64\x6b\x9c\xac\x88\x38\xab\x82\x27\xf2\xc1\xd4\xc6\x15\x1c\x9c\x77\xc7\xd6\x6f\x29\x85\xf1\x5c\x56\x2f\x39\x4d\xd9\x0a\xd8\x23\xf9\x16\xd5\x52\x7e\x25\x17\xbd\xbf\xe9\xf5\xc5\xe5\x5f\x12\x9b\x7e\xbe\x2c\x1e\x20\xc2\x48\xaa\x27\x79\x9e\x34\x19\x50\x15\x4c\x17\x19\x85\xea\x92\x41\xa5\xdd\x86\x64\x65\xfd\x56\xaf\x2d\x04\xec\xc3\x87\x0d\x1c\x94\x35\x25\xa9\xcd\xe7\x2d\x86\xa3\x7a\x94\xf7\x72\x31\x0a\x7d\xe0\x1b\xe2\xab\x5c\x0d\x84\x7f\x93\x3a\x62\x6c\x30\xa8\x7b\xb9\x4c\xd4\x27\xf9\xdf\x60\x26\x5f\x3d\x61\x94\xc6\xab\x07\xb9\x90\x8f\x17\x59\x92\x35\xed\x17\x79\x78\x6d\x37\x6d\x6e\x9b\x69\x43\xfc\xe5\x1f\xef\x1f\xe6\x57\x81\xdb\x01\x6d\x48\xdd\x0f\x9d\x35\x8a\xb7\x89\x65\x57\xdc\x63\x8b\xd4\x01\xba\xe6\x39\x68\x6a\x95\x61\xb8\xff\xd7\x62\xfb\x73\x84\x48\x6f\x46\xe1\x13\x41\x8d\xec\xf7\xf9\x2c\x63\xac\x41\x53\x37\xf1\x2d\x7b\xf8\xff\xa2\x3b\xbc\xeb\xf1\x7f\xcc\x67\x99\x4c\x4b\x04\x8c\xc3\x20\xd6\x76\x6b\x74\x8f\xce\x3a\xd0\xda\xb8\x5a\x84\xc1\xc6\xe2\x3a\xbe\x1b\x0d\x72\x75\xf1\x98\xab\x71\xe3\xcd\x67\x79\xe9\xf5\xb1\x0f\xe8\x4e\x88\xf4\x70\xb0\x63\x95\x05\xa2\x82\x3b\xd8\x95\x10\xd8\xf0\x10\x6b\x73\x40\x2d\xa2\xef\x38\x33\xba\xe0\x89\x01\x43\xbf\x2e\xb2\x5c\x9b\xb3\xd1\x4d\x5c\x03\xe0\x0a\x31\xf2\x5d\x08\x7a\xc4\x15\xa4\x0c\xe0\x4e\xa6\x59\x96\xc3\x38\x75\x7b\x2c\x05\x38\xb0\xc7\x68\x2a\x12\x01\xc1\x8a\x68\x5a\x94\x69\xa7\x9f\xd1\x59\x4e\x1d\xb8\x13\x53\x63\xb4\x46\x27\x0e\x24\xb4\xdf\x3b\xd6\x8a\x20\x1e\xf9\x74\x8f\x26\xf0\xd9\x91\x82\xf1\x98\x2b\x6d\x76\xa7\xd0\x4e\xc2\xe4\x94\x2b\x07\xe9\x20\x44\x0f\x4a\x37\x93\xf0\x3b\xa8\x51\x8c\x5b\x5e\x44\x5f\xd7\xe9\x05\x91\x8a\x76\xa5\xd8\x07\xe8\xba\x57\x94\xf0\x4e\x08\xf6\xe0\x34\x0b\x7e\xcf\x7a\xbf\x37\x45\x0b\x7e\x7f\x2a\xc2\x35\x9b\x15\xad\x16\xf7\x0f\x97\x7a\x4e\xb5\x10\x34\x4b\x5f\xa2\xb4\xbe\x7a\x9a\x54\xb1\x59\x4e\x31\xa2\x7f\xdf\xb0\x56\x94\xe2\x81\xaf\xbe\xc7\xd6\x87\x23\xdb\xa6\x3e\xcd\x55\xb3\xbc\x58\x25\xea\x94\xe9\xa9\x8b\x53\xd7\x9f\xaa\x77\x55\xca\xab\xf3\xf4\xf8\xdf\x66\x74\xce\x85\xd2\x00\xbe\x98\x4b\x3f\x9e\xaf\xcb\x65\xda\x2f\x7d\xc7\x0c\xc7\x5c\x8d\x13\xd7\x8f\xe1\xf0\xa7\xe4\xcf\x00\x00\x00\xff\xff\xa8\x05\xba\x33\xa6\x08\x00\x00")

// FileTranslationsEnGbYaml is "translations/en-gb.yaml"
var FileTranslationsEnGbYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x95\xc1\x6e\xdb\x3c\x0c\xc7\xef\x79\x0a\x36\xdf\xe1\x3b\x2d\x0f\x50\x0c\x03\xb6\xa4\x18\xba\xb5\x68\xb1\xb6\x03\x76\x5a\x19\x9b\xb1\xb5\xca\xa2\x21\x4a\xee\x8a\xc0\xef\x3e\x48\x4e\xb2\xb6\x92\xd0\x15\xd8\x72\x30\x10\x5a\xff\x9f\x48\xfe\x2d\xea\x0d\xa8\xfa\x18\x2a\xee\x3a\x34\xf5\x77\x5c\xb3\x77\xd3\x13\x60\x06\xe0\x2c\x1a\xd1\xe8\x14\x9b\x63\x98\xbf\x8f\xf1\xed\x76\x71\x49\x56\xd8\x8c\xe3\x7c\x36\xcb\xe9\xeb\x81\xac\x28\x36\x19\xc4\x4a\x49\xc5\xb6\xfe\xc8\xf0\x75\x5a\x53\x40\x34\x65\xc4\x8b\xda\xbb\xb2\xf6\xb3\x72\x8e\xda\x0f\xec\x5e\x84\x74\xd4\x79\xa1\x3a\xc3\x38\xa7\x8e\xed\x03\xdc\x08\xd5\x05\xad\x90\x0d\x29\x24\xca\xab\x29\x5e\x50\xb9\x16\xcd\x9d\xa4\xfb\x5d\x4f\xf1\x0d\x5b\xf0\xa2\x4c\x03\x8f\x8a\x38\x2a\xb0\x7c\xef\x54\x47\x09\xea\x26\x86\x0b\x9a\x96\x74\x9f\x36\x5b\x0d\x24\xa0\xcc\x86\x6d\x17\x63\x30\x7d\x1c\xae\x25\x58\xb3\x4b\x50\x01\x12\x1f\x69\x1d\x97\x96\x07\x55\x93\xec\xd7\x42\x58\x36\x9f\x01\x3c\x43\x68\x34\x8d\xc7\x86\x26\x4c\xf8\xa5\x75\x04\x67\x1c\x43\xd5\xa2\x69\x62\x22\xff\x0b\xec\x75\x49\x4e\x07\xa0\xf8\xaa\x22\x49\x7d\x39\xdb\x2d\x80\x16\x05\xd6\x44\x06\x84\xd2\xd2\x76\x1f\x26\xda\x35\x36\x69\x6b\x77\x71\xa8\x58\x6b\xaa\x5c\xe6\xdb\xa0\xaa\xe5\x7c\x93\xbf\x50\x4f\xe8\x04\xee\x5b\x74\xf0\xc0\x1e\x04\x1f\x12\xb9\x17\xb2\xc1\x87\x97\x7c\xda\x19\x84\x10\x04\x09\xa5\x63\x57\xe7\x09\xe7\x24\x12\x0a\xe0\x4d\x34\xb7\xce\xa4\x10\xc5\x7c\x6f\xc8\x92\xb5\x6c\x13\xc4\x37\xf6\x60\x68\xb2\x66\x4d\x11\x13\x57\x87\xff\x5e\x42\x40\x1d\xdc\x5f\xe4\xe9\xc1\x51\x43\x3a\xcd\xee\xe2\x7a\x05\xbb\x97\x4f\x6c\x0a\xec\xb7\xff\x6d\xb7\x8b\xe5\xf4\xf2\x74\x35\x8e\xef\x0a\x70\x21\x97\x07\x3f\x06\x16\xb4\xde\x08\xb9\x7b\xb4\x46\x99\xa6\x0c\x31\xec\x7e\x83\xc2\x88\x80\xdb\x30\x2e\x2d\x6d\xd4\xcf\x71\x0c\xa0\xf0\xe6\x36\x24\x1d\x72\x57\x99\xdd\x7a\x65\x9a\xbc\x43\x4b\xd4\x95\xd7\xe8\x48\x20\x1c\x63\x70\x78\x47\x66\x42\x99\x1a\xba\xc9\xbf\x3c\xcf\x92\x78\x9d\x16\x7f\xa9\x4c\x73\x04\xd7\x45\xd8\x71\x4c\xff\x44\x63\x2f\x54\x8f\xe3\x6d\x9a\x6d\xc3\x5c\xaf\xb9\x30\x3c\x96\xa8\x35\x74\x04\x08\x61\x59\x76\x5e\xec\xf5\xd3\xf0\x7b\x46\x98\x85\xc3\xcf\x86\x0e\x43\xf0\x28\x5c\x3e\x4b\xf6\xc6\x8d\x23\xf4\xf1\x12\x8a\x6d\xaf\x50\x6b\xaa\x9f\x6e\xb5\x98\x4f\x72\xd7\x92\x2d\x01\xb8\xd7\xe1\xd0\x0f\x54\x22\x3c\xcb\xf6\x87\xd0\x80\x3a\x5f\xec\xc9\x80\xda\x47\x73\x3e\xe1\x80\x57\x95\x55\x7d\x5a\xee\x0e\xa0\x4c\xef\x53\x37\x4e\x43\xb4\x24\x61\xef\x72\x9a\x8b\x18\x2e\x8a\xfe\xd5\x59\xd5\x1e\xff\xa4\x13\x67\x1e\x8b\xd2\xd7\xf4\x60\xaf\x79\x5d\x13\xf6\xaa\x9a\x4d\x3a\xaf\x57\x6c\x28\xbd\x3d\x0f\x1b\xfd\xad\xc6\xfd\x0a\x00\x00\xff\xff\x83\x5d\x31\xca\x5f\x09\x00\x00")

// FileTranslationsEsEsYaml is "translations/es-es.yaml"
var FileTranslationsEsEsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\xd2\x4f\xaa\xdb\x40\x0c\x06\xf0\xbd\x4f\xa1\x66\xdf\x1e\x20\xbb\x86\x42\x16\xa5\x10\xe8\x9f\x6d\xf9\x3c\xa3\xda\xd3\x78\x46\x46\x9a\x09\x94\x90\xc3\xf4\x0c\xbd\x41\xdf\xc5\x1e\xf3\x42\x1e\xe1\xd9\x03\xd9\x08\x8c\xf9\x7e\x96\x2c\xbd\xa7\xe0\xb7\xe4\x24\x46\x24\xff\x13\xbd\x94\x7c\xad\x44\x1d\x51\x56\x24\x9b\x90\x83\xa4\x2d\x6d\x3e\x3a\x56\x07\xf2\x4c\xe7\xf3\x87\x03\xab\x49\xba\x5c\x36\x5d\xb7\x66\xf8\x13\xab\x05\x49\x2b\xcc\x8f\xfa\xe6\xe9\x5f\xaa\xd0\xa7\x60\x4e\xd4\xef\xa5\xc1\x0c\x8f\x31\xcd\xfc\xf1\xb1\xfc\xe7\x90\x33\x8f\x3b\xc9\x4d\x28\x72\x2c\xc6\x7e\xc5\xf9\xc2\x51\x34\x80\xbe\x1b\x3c\x1a\x69\x63\xad\x8d\x2c\xb2\x5f\x59\x4f\xc1\x8b\xb2\x35\x82\x79\x44\x3a\xda\xf2\xa3\xff\xff\xee\x15\x2e\xc0\x68\x16\xa5\x62\xd0\xfb\x19\xde\x35\xb4\x32\xe7\x10\x79\x81\x7d\x0b\x1c\x67\x21\xb8\x1c\x4e\xad\xf9\x47\x9e\xe6\x45\xf0\xa0\x32\x8b\xba\x20\x09\x14\xd2\x2f\xd1\x88\xfa\x40\xb8\x1d\xca\x44\xbd\xe4\x85\x58\xad\x97\xb2\x9c\xeb\x5e\xc4\x9f\xe2\x41\x4e\x12\x4d\x62\x35\x8c\xe4\xc5\x36\x1d\xd1\x1b\x6f\x42\x1a\x0a\x06\xbe\x9a\xb4\xb2\xa3\xba\x1b\xa1\x19\x0a\x72\x88\x7d\x80\xd2\xc4\x35\xf4\x9b\x9b\x5d\xbe\xaa\x56\x9c\x63\x5b\x5b\x1e\x8d\x20\x58\x18\x52\xd5\x6f\x60\xeb\x92\xa1\x3d\x86\xe5\xcf\xdf\xc1\x8a\x82\x94\x9d\x4c\xec\x72\xbd\xa1\xe7\x00\x00\x00\xff\xff\xb2\xf6\x20\xd2\x94\x03\x00\x00")

// FileTranslationsFrFrYaml is "translations/fr-fr.yaml"
var FileTranslationsFrFrYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\xd2\xc1\x6a\xe3\x30\x10\x06\xe0\xbb\x9f\x62\x36\x97\x9c\x36\x0f\x90\xe3\x12\x76\x0f\xbb\x0b\x85\xd2\x5e\x8b\x22\x4d\x6c\x11\x4b\x63\x66\xa4\x40\x09\x81\x3e\x4a\x8f\xf5\x73\xe8\xc5\x8a\xe2\xd4\x0d\xb1\x05\xb9\xe8\x20\xf1\x7f\x92\x66\xe6\x27\x58\xb3\x06\x4d\xce\x29\x6f\x5e\xd4\x96\x62\x18\x56\x80\x0a\x20\xb0\xf2\xd2\xaa\x60\xc9\xaf\x61\x91\xde\xa0\x63\xea\x48\xc0\x20\x1c\x8f\xab\x07\x64\x21\x7f\x3a\x2d\xaa\x6a\x8e\x31\x07\x64\xb1\xe4\x67\xa4\xe7\xcb\x89\x41\xd8\x58\xd1\xc4\xe6\x0f\x15\x94\xfa\x2e\xa5\x18\xdf\xdf\x11\x8f\xf0\xd7\x86\x80\xcd\x2f\x0a\x45\xc7\xa1\x8b\x82\x66\x86\xf9\x9f\x7a\x47\x96\x11\x62\xb0\xad\x95\xd4\x63\x81\x10\xe4\xfc\x98\x09\xf0\x98\xf7\x23\x4b\x21\x16\x1a\xe5\xf7\x32\x73\x2f\xb2\xb6\x60\x96\xc3\xb5\xc8\xd0\xe2\xf5\x3f\xe0\x47\xc1\x8b\x5d\xb0\x0e\x27\xdc\x26\x72\xea\x31\xd7\x72\x47\x5e\xe7\x4d\x8f\x0e\x7d\x28\x28\x0d\xb6\xdd\xc4\xf8\x4d\x91\xbd\x0d\x60\x50\xc0\xfa\x1d\xb1\x3b\x9f\x08\xa4\xf7\x71\x72\x22\x6c\x69\x8a\x66\xee\xbc\x4c\xff\xf9\x8d\x42\xbb\x54\xd6\xe0\xb5\x86\xf2\x45\xa0\x2c\x2a\x80\x1b\xb6\x55\xbe\x8e\xaa\xc6\x81\x86\x99\xee\x3d\x5d\x9a\x06\x1d\x45\x06\xdd\x28\x5f\xe7\x52\x2a\x38\x47\xb1\xf4\xde\x11\x96\xa8\x35\xca\xb4\xa9\xff\x46\x41\x41\xea\x43\xea\x07\x3b\x97\x58\x1d\x50\x43\xce\xa5\x0f\x59\x8d\xf2\xed\xd8\x2b\xde\xaa\x7a\xda\xa6\x71\xda\x3c\x86\x40\xaf\x79\xda\x3e\x03\x00\x00\xff\xff\x79\x36\xca\xd4\xc3\x03\x00\x00")

// FileTranslationsNbNoYaml is "translations/nb-no.yaml"
var FileTranslationsNbNoYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x95\xc1\x4e\x1b\x31\x10\x86\xef\x79\x8a\x81\x9e\xc9\x03\xa0\xaa\x07\x08\xad\x28\xad\x40\x82\x7b\x99\xc4\xc3\xae\xb3\xf6\x78\x35\xb6\x43\xa5\x68\x1f\x87\x3e\x43\xef\x79\xb1\xca\x4e\x02\xb4\x5e\x43\x91\xda\x4b\x0e\xab\x7c\x9f\xc7\xff\x8c\xed\x23\xd0\xea\x18\x16\xce\x5a\x64\xf5\x0d\xe7\x2e\x86\xed\xef\x04\x20\x08\xb2\x37\x18\xb4\xe3\x63\x38\xbc\xb4\xb0\x5e\x4f\xaf\x48\xbc\xe3\x61\x38\x9c\x4c\xc6\x50\xb5\x22\xf1\xda\x71\x41\xcf\xb4\x5f\x38\x51\x9f\xdc\x51\xfa\xc7\xd2\x71\x45\xd0\xd4\x04\xaf\x92\x5d\x8d\xbc\xd0\x21\x50\x7b\xe2\xc2\xab\x0a\x4b\x36\x7a\x52\x85\xe1\xab\x66\x26\xd0\x30\x97\xd8\x55\x50\x4f\x92\xdc\x05\x7a\x9d\xbf\x53\x85\x0a\x2d\x72\x57\x42\x37\xd8\x75\x70\xe7\x04\x30\x80\x8a\x79\x59\x12\x78\xb6\x8f\x83\x8a\x2f\xf6\x41\x5b\x2a\x5b\xd7\xf7\x14\xb4\xaa\x40\x2d\x99\xbe\x8c\x5b\x0b\x68\xbe\x73\x62\x31\x05\x06\xce\xc2\xdc\x05\x2a\x93\x4b\x74\x5d\xd1\x2e\xc9\xf4\x60\x49\x41\xb7\xfd\xbf\x23\x29\x14\x06\xb9\x89\xd8\xd0\xb8\xe6\x24\xed\xde\xe7\x38\x36\x0f\x40\xac\x84\xb6\xa5\x78\xf0\xbd\x6c\x1e\xca\x8e\x3c\xfa\x7c\x5c\x2c\xc8\x8f\x34\x25\x73\xd0\xa2\xc0\xdc\xe8\x10\xc0\x63\x08\xb5\x81\x44\x99\x63\x53\x66\xea\x37\x3f\xfb\x9e\x0c\x78\xb4\x86\x02\x68\x2e\xa3\xa1\x45\xeb\x2a\xd1\x2c\x89\x03\x0a\x28\xca\x0d\xf6\x7a\x24\x95\xe8\x49\x52\x07\x5e\xee\x4f\x6a\x0c\xf1\x6e\x44\x0a\x87\x75\x41\x8d\xf3\x33\x6c\x52\x82\x96\x8c\xd2\xdc\x8c\x83\xee\x9e\x49\x48\xc4\x49\x89\x47\xb0\x9b\x07\x58\x6d\x7e\x08\x01\x69\x12\xe2\x7d\x87\x72\x29\xa0\x28\x1d\x99\xc7\xa6\xf3\x74\x7c\x89\x45\x8b\xcc\x64\xca\x33\x8b\x8c\x66\xe7\x54\xbf\x95\x0a\x73\x43\xb9\x5f\x10\xb4\x81\xf7\xef\xd6\xeb\xe9\xe9\x56\x72\x3e\x1b\x86\x0f\x95\x75\x3c\x95\x17\xda\x6c\xdc\x5b\x31\x44\xf6\x14\xee\x51\x58\x73\xf3\x9a\x2a\x0d\x96\xee\x3a\x7a\x36\x5d\x53\x48\x73\x0c\xb7\xe9\x0a\x15\xba\xd3\xdf\x87\x21\x69\xc1\x53\xb8\xdd\x27\xe7\x29\x84\x9c\x5c\x59\x42\xaf\xb9\xa9\x1c\x0f\x12\x6a\x98\x04\x82\x56\xc4\x79\xa2\xd2\x64\x65\x1d\x2b\x4a\xc3\x51\x6b\x72\x96\x0a\xf9\x68\xca\x70\xae\x34\x37\x07\x70\xf3\xe4\x74\xdd\x53\x99\xc9\xbb\x93\x1e\xe7\x2d\x9d\x19\xec\x3d\xa9\x61\xb8\x2d\x4b\x6f\x9c\x53\x73\x57\xb9\xeb\x26\x00\x00\x8e\x69\x77\xe9\x1d\xa4\x27\xe6\xd4\x45\x0e\xc3\x00\x7d\x7e\x6a\x72\x98\x1d\x9a\x00\x96\x9a\xb4\x9b\xc6\xa9\x74\xfe\xa7\x87\x5b\x36\xb4\x24\x75\x9a\xe4\x05\xbe\x52\xe8\x68\xcc\x17\x68\xcc\x1f\x86\x42\xb0\xf4\xb4\x42\x33\xce\x9f\xad\xd0\xc4\x74\x9a\xe0\x33\xae\xf0\x7a\x21\xba\xaf\x0a\x34\xf7\x23\xef\xef\x79\xfa\x5a\x43\x5c\x0c\x63\xcc\x65\xfe\x5c\x85\xfe\xeb\xf9\x36\x11\xff\x26\x8e\x2f\x11\xab\xe8\x5b\x82\xd8\x33\x6f\x4b\x62\x4f\x29\xc7\xe5\x15\xff\x91\x44\xe9\xa6\x7c\x6b\x1f\x97\xfa\xb7\xf9\xfd\x0a\x00\x00\xff\xff\x73\xf5\x18\x04\x8b\x09\x00\x00")

// FileTranslationsNnNoYaml is "translations/nn-no.yaml"
var FileTranslationsNnNoYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x95\xcf\x4e\x23\x39\x10\xc6\xef\x79\x8a\x82\x3d\x93\x07\x40\xab\x3d\x40\xd8\x15\xcb\x8c\x40\x82\xfb\x50\x89\x0b\xb7\xd3\xee\x72\xab\x6c\x87\x91\xa2\x7e\x1c\xde\x61\xee\x79\xb1\x91\xdd\x09\xcc\xe0\x36\x0c\xd2\xcc\x25\x07\xa7\xbf\x5f\x95\xbf\xfa\xe3\x13\x30\xea\x14\x56\xae\xeb\x90\xd5\x17\x5c\xba\x18\xc6\xdf\x19\x40\x10\x64\x6f\x31\x18\xc7\xa7\x70\x7c\xdd\xc1\x76\x3b\xbf\x21\xf1\x8e\x87\xe1\x78\x36\x9b\x92\xaa\x0d\x89\x37\x8e\x0b\xf5\xc2\xf8\x95\x13\xf5\x9f\x3b\x49\x5f\xac\x1d\x57\x00\xba\x06\x78\x57\xd9\xd6\x94\x57\x26\x04\x6a\xce\x5c\x78\x17\xd1\x51\x17\x3d\xa9\x82\xf0\xd9\x30\x13\x78\xd7\xc1\x52\x62\x8b\x3e\x54\xf4\x9e\x24\x05\x28\xf4\xb7\xf9\x1c\xa5\x22\x0b\x0d\x72\x5b\xaa\xee\xb0\x6d\xe1\xc1\x09\x60\x00\x15\x73\x68\x12\xf8\xe1\x36\x47\x15\x5e\xec\x83\xe9\xa8\x2c\x60\xdf\x53\x30\xaa\x22\x6a\xc8\xf6\xa5\xe9\x46\xc0\xf0\x83\x93\x0e\x93\x6d\x90\x1c\x70\x81\x4a\xff\x92\xba\x8e\x68\xd6\x64\x7b\xe8\x48\x41\x3b\x7e\xef\x26\xbc\xb0\xc8\x3a\xa2\xa6\x69\xcc\xd9\x68\x7c\xf6\x63\xf7\x04\xc4\x4a\x08\x7c\x2f\xbb\xa7\x96\x02\x04\x63\x2b\x89\x3d\x53\x7d\x5c\xad\xc8\x4f\xd4\xe6\xc0\x68\x50\x60\x69\x4d\x08\xe0\x29\xd4\x2a\xac\x51\x96\xa8\x4b\x6f\xfd\xee\x5b\xdf\x93\x05\x8f\x9d\xa5\x00\x86\xcb\x4c\x68\xd5\xb8\x8a\x45\x6b\xe2\x40\x2d\x28\xca\x85\xf6\x64\xa8\xb4\x27\x7a\x92\x54\x8a\xb7\x0b\x95\x2a\x44\x86\xc7\x3e\x2d\x21\x9d\x0b\x6a\x1a\xb0\x40\x4d\xec\xa1\x23\xab\x0c\xeb\x69\xa1\x7b\x64\x12\x12\x71\x52\xca\x23\x74\xbb\x27\xd8\x90\x10\x90\xd1\x28\xc4\x87\x52\xe5\xbe\x05\x45\x69\x84\x0e\xe5\x27\x9e\x4f\x87\x58\x35\xc8\x4c\xb6\x9c\x61\x64\xb4\x7b\xa6\xfa\x29\x55\xd8\xa0\x8c\x25\xcb\x5d\xf0\xf7\x5f\xdb\xed\xfc\x7c\xa4\x5c\x2e\x86\xe1\x9f\x4a\x20\x4f\xe5\x86\x5b\x54\xc0\x15\x44\x64\x4f\xe1\x11\x85\x0d\xeb\xf7\x58\x24\x60\xda\x76\x4d\x23\x0f\x52\x37\xc3\x7d\x5a\xa7\x42\x0f\xe6\xeb\x30\x24\x60\xfa\xef\xfe\xe0\x9a\xcf\x1d\x39\x61\x53\x6f\x58\x57\x46\x84\x84\x5a\x26\x81\x60\x14\x71\xee\xa6\xd4\x55\x19\xc6\x8a\x72\x5f\xd4\xea\x9b\xa9\x42\x3e\xda\xd2\x96\x1b\xc3\xfa\x08\xee\x5e\xa0\xae\x7d\xc9\x32\x81\xf7\xd0\xd3\x7c\xa3\x0b\x8b\xbd\x27\x35\x0c\xf7\x65\xee\xda\x39\xb5\x74\x95\x5d\x73\x85\xd6\x42\x47\x3a\xe7\xa9\x9d\x4a\x03\x5d\x25\x4c\xae\xcc\x19\x00\x80\x63\xda\xef\xce\xa3\xf4\x5e\x9d\xbb\xc8\x61\x18\xa0\xcf\xef\x56\x1e\xf2\x16\xad\xc5\xd7\x91\xe6\xc7\xa3\x3a\x34\x24\x75\x3d\xca\x9b\x84\x57\xc9\xae\x3d\x6d\xd0\x4e\xdf\xf6\x62\x83\x36\xa6\x69\x82\xff\x71\x83\xb7\x2b\x31\x7d\x79\xdb\x3d\xc0\x70\x3f\xf1\x1e\x5f\xa6\xd3\x9a\xc4\xc5\x30\xa5\xb9\xce\xc7\x55\xd1\x1f\x9d\x6f\x1b\xf1\x57\xec\xf8\x14\xb1\x2a\xfd\x88\x11\x07\xcd\xc7\x9c\x38\xa8\x94\xe3\x72\xc9\xff\x4b\xa2\x8c\x2e\x5f\xdd\xe7\x50\xbf\xd7\xbf\xef\x01\x00\x00\xff\xff\x09\xd9\xf3\x8c\x9b\x09\x00\x00")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}



var err error



  
  err = FS.Mkdir(CTX, "translations/", 0777)
  if err != nil && err != os.ErrExist {
    log.Fatal(err)
  }
  

  
  err = FS.Mkdir(CTX, "assets/", 0777)
  if err != nil && err != os.ErrExist {
    log.Fatal(err)
  }
  

  




  
  var f webdav.File
  

  
  
  var rb *bytes.Reader
  var r *gzip.Reader
  
  

  
  
  
  rb = bytes.NewReader(FileAssetsGraphsJs)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "assets/graphs.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileAssetsStyleCSS)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "assets/style.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileIndexHTML)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileTranslationsEnGbYaml)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "translations/en-gb.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileTranslationsEsEsYaml)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "translations/es-es.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileTranslationsFrFrYaml)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "translations/fr-fr.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileTranslationsNbNoYaml)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "translations/nb-no.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(FileTranslationsNnNoYaml)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "translations/nn-no.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  


  Handler = &webdav.Handler{
    FileSystem: FS,
    LockSystem: webdav.NewMemLS(),
  }


}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

  // If the buffer overflows, we will get bytes.ErrTooLarge.
  // Return that as an error. Any other panic remains.
  defer func() {
    e := recover()
    if e == nil {
      return
    }
    if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
      err = panicErr
    } else {
      panic(e)
    }
  }()
  _, err = buf.ReadFrom(f)
  return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
  if err != nil {
    return err
  }
  n, err := f.Write(data)
  if err == nil && n < len(data) {
    err = io.ErrShortWrite
  }
  if err1 := f.Close(); err == nil {
    err = err1
  }
  return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}


