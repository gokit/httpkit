package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "http-api-json.tml",
        
          "http-api-mock-functions.tml",
        
          "http-api-mock.tml",
        
          "http-api-readme.tml",
        
          "http-api-test.tml",
        
          "http-api.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "http-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xc1\x6b\xc2\x30\x14\x87\xcf\x0d\xe4\x7f\xf8\xad\x87\x61\x41\xea\x7d\xc3\xc3\x18\xec\xb0\x83\x0e\xc4\xd3\x18\x18\xe3\xab\xc6\xb5\x89\xbc\x46\x65\x84\xfc\xef\x23\xad\x32\x0f\x13\xe6\x2e\x85\xe6\x3d\xbe\xef\xe3\x1d\x14\x23\x84\xda\x1d\x89\x51\xce\x3c\xef\xb5\x2f\xa7\xcb\x2d\x69\x5f\x4e\x54\x43\xdd\x27\xc6\xd7\xd9\x74\x82\x31\x16\x21\xa0\x51\xbb\xee\xef\xb4\x8c\x7c\xdb\x3a\x9b\x23\xcf\x11\xe3\x42\x0a\x29\xfe\x46\x7c\x66\x52\x9e\x7e\xe3\xf6\x93\x27\xed\x8d\xb3\xff\xa5\xcf\x77\xab\x2b\xf4\x7e\x72\x8d\x3e\x1a\xa1\x76\x6a\x95\x56\x5f\x1c\x83\xc9\xef\xd9\xb6\x50\xb0\x74\x84\xb1\xad\x57\x56\x13\x5c\x05\x85\x10\xce\xfa\x37\xa5\x3f\xd5\x9a\x62\x2c\x7f\xde\x2e\x92\x62\x44\xc5\xae\x81\xdf\x10\x76\xec\x0e\x66\x45\x48\x5e\x68\x67\x3d\x59\x5f\x4a\x51\xed\xad\xbe\x14\x0f\x4e\x23\xb4\x9e\x8d\x5d\x17\x18\xdc\x60\x1b\x82\x98\x1d\x17\x08\x52\x64\xe9\x5c\x54\x53\x73\x4b\x6e\x3a\x44\x66\xaa\x84\xc1\xc3\xb8\x6b\x2d\xe7\xb6\x51\xdc\x6e\x54\x3d\x78\xff\x58\x7e\x79\x3a\x17\x16\x43\xdc\x27\x7e\xf1\xd8\xad\xdf\x8d\x61\x4d\xdd\x99\xb3\xfe\x78\xb7\x88\x43\xdf\x2e\x45\xd6\x37\x9c\x08\x89\x3f\x4c\x5c\x29\xa2\x14\xdf\x01\x00\x00\xff\xff\x5f\xe2\x7b\xa9\xb1\x02\x00\x00"),
          path: "http-api-json.tml",
          root: "http-api-json.tml",
        },
      
        "http-api-mock-functions.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x52\xbb\x8e\x9c\x30\x14\xed\x91\xf8\x87\xab\xad\x40\x5a\x99\x3e\xa9\xa2\x49\x93\x22\x33\x91\x92\xfd\x00\xaf\x7d\x66\x70\x16\x6c\x64\x5f\x60\x22\xc4\xbf\x47\xc6\x10\x2d\xd1\x4e\xa4\x89\xd2\xa4\xe1\x75\x39\x0f\xdf\x73\xaa\x8a\x3e\x3b\xf5\x72\x78\x22\x8d\xb3\xb1\x08\x24\xa9\x93\xea\x45\x5e\x40\x0d\x06\x34\x34\x48\x6f\xe4\x73\x03\x62\x47\x52\x29\x84\xf8\x8b\xb1\x81\xa5\x55\x20\x77\x26\x49\xed\xc2\x20\xf2\x6c\x90\x7e\xa3\x5b\x3f\xe6\x59\x9e\x55\xd5\xfa\x42\x9d\x77\x83\xd1\x8b\x48\x60\xdf\x2b\xa6\xb1\x36\xaa\x26\xe5\x2c\x4b\x63\x03\xb5\xe0\xda\xe9\x10\xb5\x94\x87\x64\x90\xb4\x9a\xfa\x4e\xc7\xc7\x69\x12\x5f\x17\x94\x38\x3d\x7f\x87\x62\x71\x94\x2d\xe6\xf9\x7d\x9e\xf1\x8f\x0e\x9b\x46\x22\x9e\xe6\x55\xf9\x69\xc1\xde\x80\xae\xcc\xd1\x0f\xae\x86\x8d\xbd\xbc\x12\xf9\x92\xd6\x30\xcf\xe2\x16\xda\x43\x39\xaf\x69\x34\x5c\x13\xd7\xa0\x8b\x19\x60\x37\xb7\x86\xd1\x8a\x3c\x3b\xf7\x56\x15\xc9\x5a\xb9\x9a\x29\x14\x5f\x97\x23\xe3\xca\xe2\x90\xee\x8f\x1b\xd9\x1d\xfa\x8f\xab\x94\x8f\xa0\x44\xfd\x41\xb1\x71\x76\x0f\xdd\x4d\x76\x04\x25\x15\x77\xc9\xc1\x7b\xe7\xcb\x29\xcf\x88\x88\xd2\xb5\xaa\xe8\xdb\xe9\xe3\xa9\xd0\xb1\x2b\xae\x83\x2f\xdf\xfd\x1a\x9c\x06\x78\x6f\x34\x28\xee\x20\xaa\xa7\x43\x5b\x0e\x69\x65\x63\x2d\x99\x42\xed\xfa\x46\x53\x2d\xbb\x0e\x56\xc4\xd0\x22\xda\x83\x7b\x6f\xef\x59\xc6\xb4\xf9\x0b\xe2\x88\xb1\x78\x38\x3a\xa6\x4f\x6d\xd7\xa0\x85\x65\xe8\x87\x32\xcf\xb6\x4a\x1c\x96\x66\xdd\x0a\x35\xf5\x2e\x56\xc2\x62\xfc\x9b\x3a\x9c\xbd\x6b\x5f\xd5\x61\xed\x71\xac\xe8\xef\x75\xf8\xa3\x91\xb7\x5b\x82\x06\x6d\x34\x95\xa0\x6f\xc5\xbd\x9b\xfc\x83\xb8\x69\xda\x42\xf9\xcf\xb2\xfe\x19\x00\x00\xff\xff\x51\xd2\x08\x10\xdb\x04\x00\x00"),
          path: "http-api-mock-functions.tml",
          root: "http-api-mock-functions.tml",
        },
      
        "http-api-mock.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcd\x6e\xdb\x38\x10\xc7\xef\x06\xfc\x0e\x93\x00\xbb\x90\x0a\x41\xbe\xb7\xc8\x21\x9b\x14\x8b\x00\xdb\xd4\x68\xd0\xd3\x62\xb1\xa0\xc5\x91\xcd\x35\x45\xaa\x24\x25\x27\x10\xf4\xee\x0b\x8a\xd4\xb7\xe3\x54\x45\x2e\x51\x2c\x92\xf3\xf1\x9f\xdf\x0c\xb5\xd9\xc0\x17\x99\x1c\x6f\xb7\x0f\x5f\x73\x54\xc4\x48\x05\x14\x53\x26\x50\x03\x11\xa0\x8d\x2a\x12\x53\x28\x84\xd3\x81\x25\x07\x60\x59\xce\x31\x43\x61\x34\x98\x03\xc2\xf0\x54\x2a\x15\xe4\x4a\x96\x8c\x32\xb1\x87\xf5\x6a\xb3\x01\x02\x99\x4c\x8e\x50\x68\xb2\x47\x60\x02\x0c\x6a\xa3\x9b\x8d\x85\x46\x38\x31\x73\x68\xac\x54\xd5\x8f\x42\x1a\x84\xf8\xa9\xf1\x16\x7f\xdd\xfd\x87\x89\x89\x1f\x49\x86\x75\x0d\x07\x63\x72\xeb\x28\x5e\xaf\xcc\x4b\x8e\xb3\x68\x5d\x88\xd5\x7a\x05\x00\xa0\x8d\x54\x08\x1f\xf4\x8b\x48\x32\x92\xc7\x5f\x48\xbe\x5e\xd5\xeb\x55\x13\xcd\x23\x9e\xa6\x67\x15\x9a\x42\x09\x0d\x04\x04\x9e\x80\x09\x6d\x88\x48\x10\x64\x0a\x64\xea\x27\x5e\xaf\xd2\x42\x24\x67\xac\x04\x21\x7c\x98\x1a\xf6\xe1\x38\xfb\xf0\xfb\x64\xd9\xaf\x76\x01\x7f\xb4\xee\x83\x41\xd0\x61\xe4\x76\xd4\x7d\xf4\xf7\xc8\xd1\xa0\x57\x18\x9d\xfc\xb2\xb1\xc7\xa4\x00\x23\x41\x61\x26\x4b\x04\x02\x7b\x56\xda\x0a\x28\x4c\xa4\xa2\xc0\x28\x0a\xc3\x52\x86\x14\x76\x2f\xf0\x70\xdf\xe6\x11\x64\xb3\xa8\x43\xef\x24\x48\xcc\x33\x24\x52\x18\x7c\x36\xf1\x9d\x7b\x46\x90\x17\x3b\xce\x92\x87\x7b\x2b\x38\x13\xfb\x10\x50\xa9\x3e\x53\x96\xc2\xbf\x11\x1c\x88\x86\x8f\x37\x90\xc5\x4d\x5a\xf1\x5f\x92\xd0\xa0\x3d\x17\x7e\x82\x2b\xbb\x3e\x48\xde\xcb\x93\x66\x26\xfe\x6c\x8d\xa5\xc1\xf5\x37\x17\x35\x95\xa8\x41\x48\x03\xf8\xcc\x2c\x33\x0d\x2b\x8c\xc2\x6f\x3f\xae\xfb\x48\xc2\x4e\x24\xf7\x4f\xeb\xd6\x67\x31\xd9\xe6\x9d\x09\xc6\x7b\x51\xff\x44\x73\xcb\xf9\x00\x03\xcd\x99\xaf\x3f\xe7\x40\x4a\xc2\x38\xd9\x71\x6c\xb5\x94\x29\x34\x08\x56\x55\x8b\xea\x96\x24\x47\xb2\xc7\xba\x8e\xfb\x77\x23\x7c\x2f\xca\xed\xdc\x9f\x97\x5b\x2a\x8a\xca\x6b\xed\x7f\xfd\xf1\xd2\xfd\xce\x5d\x47\x99\x08\x14\xea\x5c\x0a\x8d\x6a\x8b\x6a\xeb\xdf\x86\x10\xfc\xfd\xcf\x82\x20\x23\x67\xaa\x29\x68\xd8\x16\xa8\x24\xca\xe7\xad\x61\x91\xb5\x69\x39\xbe\x11\xb1\xc7\xc0\xaa\x10\x1c\x23\x28\xad\x2f\x54\x29\x49\xb0\xaa\x43\xd8\x49\xc9\x87\x48\xb0\x14\x90\x63\x16\x81\x3c\x5a\x92\xca\x38\x58\xe0\x39\xfc\x64\x8f\x0d\xac\xb9\xba\xbb\x1c\x6e\x80\xe4\x39\x0a\x1a\xf8\x17\x51\xe3\x28\xec\x37\x77\x81\x0f\x68\x31\xaa\x40\x4f\x59\xd8\x2e\xfb\xa5\xce\x0c\x47\xd1\xda\x0c\xa3\x19\x5f\x76\xbb\x62\x58\xda\x79\xda\x72\xb4\x23\x1a\x29\xd8\xb6\x3d\x74\x1d\x4d\x3b\xac\xdf\x42\xe6\x27\xdb\x73\x89\x70\x6d\xed\xbd\x74\xae\x02\xa9\x2c\x04\x7d\xbd\x9d\xbb\xbe\xbf\x6a\x36\xce\xdb\x7a\x81\xff\xaa\x8e\xde\x6d\x0a\xa8\xcf\x03\x7e\x6c\x26\xcb\x10\xea\xd3\x1a\xa3\xf4\x6b\x49\x35\xb2\xea\xf8\x11\x4f\xf3\xa4\x32\x62\x92\x43\x33\x50\xae\x67\x39\x78\xc4\x5c\x2a\x23\xa4\xbe\xe7\x94\x18\x84\xa2\x79\xe8\xd9\xb8\xef\x2e\xd5\x3d\x2b\x51\x34\xd7\x5a\x49\x78\x81\x17\xa1\x72\x36\x7f\x8a\x2b\xd7\x35\x56\x06\x77\xe8\x36\xb1\xf7\xcf\x58\x8c\xd1\xca\x58\x5d\x27\x48\xb5\x5e\x55\x95\xd5\xd8\x6f\x7d\xd0\x4f\x2c\x63\x9c\x28\xb0\xd3\x63\x38\x3b\x9e\xec\xdf\x8e\xb9\x51\xcb\x0e\xe7\x79\x55\x01\x72\x8d\xdd\x71\x21\x4f\x43\x0a\xce\x03\xbc\x8c\x0b\x78\x03\x8c\x61\xa5\x1f\xa5\x71\xdd\xb3\x61\xa2\x24\x9c\xd1\x69\x95\xed\xc3\x47\x8a\x3e\x52\x54\xca\x86\x6a\x6b\x73\xf7\xdd\xeb\xf2\x4a\x2c\xb6\x50\x51\x9f\xe3\x40\x13\x3b\x3d\x95\x82\xab\x1b\x2b\xcb\xf9\x20\x5f\xb9\x33\x9d\xd0\x3e\x9a\x78\xdb\x09\xee\xdf\x4c\x87\x5f\xa7\xba\xa0\x8d\xe8\x2d\x9d\x77\x0a\x2d\x9d\x84\xd2\xf6\xa3\xaa\xfd\x0c\x11\x46\xb6\x5c\x0e\x70\x75\xce\x2f\xb1\xe9\x2c\x9e\x67\xb3\x45\xd1\xed\x39\x87\xe2\x68\x65\x52\xd0\x5f\x1c\x90\x9e\x5c\x67\xf9\x2d\x72\x71\xac\xe6\x1c\x5f\xec\x3b\x7c\x06\xf1\x79\x34\x9c\xe3\x8b\x68\x2c\x20\x62\xf9\x3c\x7b\x27\x80\xda\xe4\x66\x24\xfd\x1f\x00\x00\xff\xff\xb5\xf8\x5b\x2b\x91\x0c\x00\x00"),
          path: "http-api-mock.tml",
          root: "http-api-mock.tml",
        },
      
        "http-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\xdf\x6f\x22\x37\x10\xc7\xdf\x91\xf8\x1f\x46\xe1\x25\x39\xe5\x36\x69\x75\x7d\x41\xea\x43\x2e\x47\xd3\x54\xed\x05\x15\xee\xe9\x74\xea\x1a\x7b\x60\x9d\x33\xb6\x6b\xcf\x92\xa0\xd5\xfe\xef\x95\xbd\x86\x40\x0a\x04\xd2\xa8\x55\xaa\xcb\x43\x56\xb2\x67\xbf\xf3\xc3\x9f\x19\x2f\x55\x95\x0d\xc8\x95\x9c\xb2\x9b\xd1\x2d\x72\xca\x3e\xb2\x29\xd6\x35\xfc\x3c\x1c\xf6\xe1\xa2\x7f\x0d\xed\xd6\x8f\xbb\xff\xda\xad\x76\xeb\x49\x15\xe9\x81\x01\x2b\xc9\xc0\x04\x35\x3a\x46\x28\xa0\x20\xb2\xc0\xac\x84\xb1\x71\x40\x05\x82\x8f\x12\x90\x6f\x54\x4b\x92\x79\x16\xfc\x0d\x0b\x8c\xb2\x78\x6f\x91\x93\x8f\x6f\x97\x1e\x1d\x90\x01\x39\xb5\x0a\xa7\xa8\x09\x98\x16\x60\x9d\x99\x49\x81\xd1\x62\xc4\xf8\x57\xd4\x02\xa4\x26\x74\x63\xc6\x31\x98\x27\x03\x01\xa5\x16\xe8\x94\xd4\x08\x62\xd4\x98\x30\x4e\xd2\xe8\x6e\x70\x98\xe7\xf9\xc4\xb4\x5b\x34\xb7\x08\xef\x1f\xcb\x54\xed\x16\x00\xc0\x07\x54\x48\x78\xcc\x8d\x26\xbc\xa7\xec\xb2\x79\x9e\x86\xb4\xa4\x9e\x9c\x00\x3a\x67\x5c\x63\x7a\x85\xb4\xdd\xee\xf8\x21\xff\x3e\xe3\x5f\xd9\x04\xeb\x3a\xdb\x52\xe1\xd3\x46\xf5\xa4\x91\xfd\x64\x05\xdb\x1e\xc1\x29\x54\x55\xd6\x98\x5c\xc4\xcc\xd6\xe5\xd7\x76\xd6\x9c\x3c\x0e\xfd\x42\xa9\xed\x3e\x16\x4f\xa9\x29\xfe\x3b\x81\xe3\xcf\x5f\x0e\xca\x28\xbe\xb9\x9a\xd6\xa5\xc3\x8d\x69\x55\x55\xd6\x6c\x6d\x4a\x67\x6d\xe7\x51\x3a\xcf\xab\x70\x1d\x31\x08\x34\x5c\x53\x20\xcf\x78\x6c\xc8\x1b\x1b\xa5\xcc\x9d\xd4\x13\x98\x22\x15\x46\xf8\x48\x34\x32\x5e\x00\x6a\x61\x8d\xd4\x14\x21\xea\x74\x52\x2a\x50\x41\xff\x66\x30\x84\xb3\x2d\x3e\xcf\xa0\x0e\xd6\x1d\xf8\x2d\xea\x75\x21\x1f\x97\x9a\xc3\x71\x68\x96\x37\xa1\xa7\x2e\xfa\xd7\x27\xcb\xb2\xd0\x3d\xbc\x09\xbd\x54\x92\x54\x8b\xe2\xa4\x23\x8b\xd1\x26\x9f\x0e\x39\xca\x59\x0a\x79\x49\xbd\x43\x6e\x9c\x00\x33\x8e\xcb\x3b\xeb\x06\x91\xfe\xbb\x42\xf2\x22\xf4\xb3\x40\x25\x71\x86\x22\xbe\xd8\x6e\xfd\x32\xb8\xf9\x08\xf1\x88\x34\x85\xbe\x0a\xab\x8b\xf6\xcf\x60\x58\x48\x0f\x77\x52\x29\x90\x1a\xa8\x74\x1a\x1c\xc6\x07\x03\x87\x3e\x34\xb1\x9c\x85\x01\xc0\xa8\xf4\xc0\x8d\xc0\xd8\xe6\x6f\xa1\x17\x1b\x1c\x05\x5c\x26\xe5\xe1\xdc\x62\x17\x52\x4b\x86\xac\x13\x22\xcd\xf6\xdb\x66\x9b\x59\xab\x24\x67\x21\x85\xb3\x5b\x6f\xf4\xf2\xe0\x56\x04\x7f\xc7\x3f\x4b\xf4\x04\x7d\xe6\xd8\x14\x09\x9d\x4f\xa2\xbb\x8c\xdf\x1b\x31\x4f\x66\x8d\x6e\x55\xc1\x94\xd9\x98\xfb\x5a\xe5\xe0\x68\xe4\x8d\x3e\x82\xa3\xdb\xf8\xa8\xeb\x4d\xaa\x83\x26\xdd\x4b\x23\x70\xe9\xfb\x27\x26\x55\xe9\xb0\x0b\x3f\x9c\x9f\xb7\x5b\x83\x92\x73\xf4\xbe\x0b\xdf\x9f\x7f\xb7\x39\x2e\x6f\x8d\xf6\xb8\x3b\xb0\x06\xb1\xed\x21\x75\x3a\x70\xd5\xdb\x41\x63\xd7\x96\x23\x25\xf9\x1f\x52\xec\x41\x65\x1c\x6d\x4f\x21\x79\x85\x14\x00\x70\x91\x47\x06\x13\x39\x0b\xdd\xf3\x37\x16\x37\xde\x29\x91\xc2\xb1\x33\xd3\x35\xc4\x12\x4f\x49\x26\x80\x1e\xd8\xf6\xa5\x22\x60\xc1\x45\xa8\x44\xbb\xe5\x52\xbd\x32\xb8\xa6\x70\x61\x3c\x6a\x86\xfc\x21\xd3\x1c\xec\x02\x8c\x20\x10\xcd\xc2\x02\xa5\x4b\x46\xa0\x26\x39\x9e\xc7\x8d\x26\xee\xff\x92\xd8\x20\xb8\x7a\x4a\x07\x10\xfc\x0f\xb1\x3c\x7f\x3e\x96\xfb\xf5\xcb\x53\x70\xee\x87\x64\xbc\xb2\x0e\xa3\x52\xa9\x74\xb0\xfe\x79\x44\xbe\x92\x09\xf6\xef\x9e\xff\xe7\xbd\x01\xf8\xb2\x4a\x40\xff\xd3\x4b\x8d\xa7\xc5\x27\xd2\x53\x2c\x34\x76\xc0\x88\x70\x6a\xc3\x17\xa6\x81\x32\x2d\xfd\x3f\xc6\x15\xdc\x49\x2a\xd6\xf5\xe2\xa9\xb8\x04\xca\xc8\x88\xf9\x2b\x1d\x69\xab\x8c\xad\x7e\xd5\xbe\xfc\xa5\xfc\x6e\xa9\xb0\x27\xff\x0f\x48\x7f\xe8\xfd\xda\x1b\xf6\x5e\x88\xea\xc5\x4f\x8f\x7d\x26\x9c\x88\xb6\xdf\xae\xdd\x57\x7d\xed\xbe\x3b\x6c\xec\xe6\x79\xfe\x57\x00\x00\x00\xff\xff\x14\x8a\x04\x80\xe9\x0f\x00\x00"),
          path: "http-api-readme.tml",
          root: "http-api-readme.tml",
        },
      
        "http-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4f\x4f\x23\xc7\x13\x3d\xff\x90\xf8\x0e\xf5\x1b\x89\xd5\xcc\xc6\x6a\x27\x57\x47\x3e\x10\xfe\xec\x92\x10\xb0\xb0\x51\x8e\x51\x7b\xa6\x0c\xbd\xdb\x9e\x36\xdd\x35\x06\x64\xf9\xbb\x47\xd5\x3d\x36\xe0\xb5\xc1\xc6\x98\xec\xc6\xb3\x87\x45\x8c\x67\x5e\x55\xbf\x7e\xf5\x6a\xaa\xcd\x50\x5a\x88\x77\x77\x00\x00\x70\x88\x39\x39\x68\x42\x1f\xc9\xaa\xd4\x89\x33\xbc\x8d\xd3\xc2\x91\xe9\x8b\x36\xc9\xf4\xeb\xa1\x72\x03\x2d\xef\x63\xe3\x44\x9b\x32\x53\x50\x92\x84\x27\x87\x68\x9d\x32\x39\x34\x21\x1a\xfe\x12\xed\xee\x24\xbb\x3b\xbb\x3b\xbd\x22\x4f\xc1\xe2\x95\x72\x84\xf6\xc2\x14\x84\xb1\xe5\xff\x2d\x5c\x13\x0d\x0a\x52\x5a\xf8\xab\xb6\x06\x72\xa0\xe0\x23\x5f\x95\x03\x25\x3e\x77\x3a\xad\xfd\xd6\x49\x6d\x8a\xea\xc8\xaa\xfc\xaa\x06\x16\x9d\x29\x6c\x8a\xe5\x85\x04\x46\xbb\x3b\xff\x0b\x90\xe2\xb3\xcc\x33\x8d\x71\xf4\xe9\xa8\x13\xd5\x7a\x7d\x12\xed\x81\x55\x39\xf5\xe2\xa8\xbe\xe7\xea\x7b\x2e\x9a\xc2\x3d\xe0\x24\x3e\xb0\xf8\x84\xb4\xaf\x75\xb2\x3c\x56\xbd\x31\x28\xba\x5a\xa5\x7f\xab\xec\x79\x58\x4f\xc3\x2c\x6a\xeb\xbc\xbd\x72\x8a\x07\x16\x25\xe1\x7c\xb8\xcb\x75\x92\xbc\x1c\x64\x01\x78\x16\xf6\xf0\xe8\xf4\xa8\x73\xb4\x06\xf2\x21\x6a\xf4\xc8\x63\xce\xba\x5e\x87\x0e\x3a\x0a\x54\x8f\x46\xa2\x4d\xb6\x48\x49\x9c\x77\xbf\x60\x4a\xe2\x4c\xf6\x71\x3c\x86\xa1\xd4\x8a\xd3\x71\x40\xd7\x08\x96\x45\x88\x43\xa9\xc1\xf4\x40\xc2\x82\x87\x3c\xb4\xc5\xd4\xd8\x0c\x7a\xd6\xf4\x61\xa2\xa2\x52\x7f\x2f\x46\x8d\x09\x3e\x12\x3a\x52\xf9\x95\xe8\x24\x23\x4e\x96\x15\x9d\x75\xa1\xd1\x84\x33\xbc\xfd\xd3\xa4\x5f\xf7\x5b\x27\xe7\x03\xb4\x92\x8c\x8d\xc3\x26\xb0\x60\x1b\xcd\x69\x30\x2e\x94\x50\x3d\x35\xc8\xba\x4c\x27\x59\xc4\xc9\x1d\x5e\xea\x0c\x55\xdc\x05\xc1\xc7\xb9\xd2\x01\xa7\x5e\x87\x8b\xb2\x42\xc0\xef\x80\x83\x5b\x45\xd7\x50\x16\xca\x95\x35\xc5\x40\xf0\xee\x94\x37\xc5\x8c\xeb\x09\x7e\xc4\x7c\x34\x1a\x91\x39\x35\xb7\x68\x61\xee\x12\xa3\x64\xb2\x2a\xd4\xd8\xaf\x01\x5a\xcb\xa9\x69\x23\xb3\xdf\xdb\xe7\x67\xc7\xc6\xc6\xa3\x91\x5e\xf4\x7c\x09\x12\x24\xc8\xf7\x97\x35\xaf\x7a\x1e\xe8\xff\x4d\xc8\x95\xe6\x52\x84\xf2\x1f\xb3\xe9\xc4\xb1\x54\x1a\xb3\x38\x6a\x5f\x9b\x42\x67\x70\x2d\x87\x08\xae\x48\x53\x74\xae\x57\x68\x7d\xef\xc3\x63\x06\x8c\xd8\x80\x55\xe2\x43\x03\xf6\x7e\xba\x11\x91\x5f\x48\x99\xcc\x38\xfc\x08\xa1\x5b\xd2\xb9\xcd\x84\x16\x0f\x54\xa6\x74\xc7\x24\xa6\x26\x27\xbc\x23\xaf\x80\x39\x34\x37\x21\xeb\x96\xd5\x1b\xa7\x74\x57\xf3\x1f\xbd\x11\x81\x4e\x0e\x31\x5b\x54\x17\x93\x9a\x58\x9b\xab\x65\xa2\x3c\xa2\xc5\xe2\xcd\x54\x60\xac\x7d\x26\xe6\x02\x6f\x0a\x74\x54\x1a\x2a\x7c\x63\x29\x2f\xc9\x77\x2a\xf5\xa4\x06\xa1\x72\xde\x82\xbe\xd4\xef\x4a\xc6\x19\x73\x76\x1b\x27\x72\xb5\x78\x4f\x28\x75\x13\x36\x39\x54\x60\x94\xef\x41\xfb\x20\x39\xf6\x05\xd1\x46\x3b\x44\x6e\x9f\xb1\x45\xc7\x8e\x7c\x33\xfd\x5c\xf5\x18\x47\x1c\x98\x0c\x99\x32\xbf\x35\x6d\x92\x54\xb8\xf3\x3f\x96\x23\xcf\x62\x8a\x8a\xb5\xe0\xfc\x63\x90\x32\xd4\x5e\xe9\xba\x16\xdd\xc0\xe4\x0e\xa1\x5b\x10\x5c\x19\x82\xbd\x4c\x44\xb5\x27\x51\x6a\xd3\x04\x96\xa5\x6f\xa9\x88\xbc\x25\x4f\xe2\xcc\x2e\xf9\x37\x93\xdd\x8b\x53\xcc\xe3\x04\x9a\x4d\xf8\xf9\x15\x4a\x99\xe6\xf1\xb0\x4a\xc6\x8c\x5e\x25\x83\x67\xc0\x66\x3a\xe5\x6b\xda\xa4\xd6\x4b\x37\x4a\xb9\xa0\x55\x2e\xdb\x27\xbf\xe3\x2e\x19\x5e\x36\xab\x56\x59\xb5\xca\x2d\x6b\x95\x33\x53\x04\x33\x28\x5a\xfe\x65\xfd\xe4\xb0\x6a\x9e\x55\xf3\xdc\x92\xe6\xb9\x40\x1d\xc1\x5c\x66\xfa\xa7\x57\x96\x32\xf9\xd2\x53\xa6\xef\x3c\x73\x9b\xe7\xb3\x61\x7f\xac\x39\xf3\x2d\x3a\xe8\x50\x5a\xbf\x3d\xd0\xbd\x27\x96\x52\xd1\xeb\xa1\x0d\x1f\xf9\x5d\xfb\xcb\x2a\xc2\xb6\x3f\xca\x59\xb5\xa9\x2e\x61\xa0\xe1\x98\x65\xbd\x61\xe3\x03\xe7\x59\x39\xe6\x9b\x3a\xe6\x41\x99\xd7\x3b\xd8\x66\x19\xea\x7d\xbc\xb3\x0c\xf6\x9f\x36\xd0\x70\x5a\x38\x63\xa0\x45\xb8\xb8\x31\xfb\x0c\x41\xb7\x6d\xfc\xa8\xcc\xb3\x32\xcf\xca\x3c\x7f\x18\xf3\x7c\xe1\xcc\x60\xba\x98\xb2\x62\x93\x4d\x1c\x0c\xbc\xe1\xa8\xff\x68\x4d\xa3\x11\x58\x99\x67\xc7\x0a\x75\x36\xa9\x7a\x88\x78\xa5\x11\x44\x5f\x9c\xc9\x23\x88\x1e\xbe\x0f\x82\xf1\xf8\x89\x83\x15\x33\xfe\xf5\x64\xd9\x8d\x26\x30\x00\x17\xc9\x51\xce\x32\xb1\xf1\x87\x6e\x91\x88\xf0\x4b\xec\xe7\xff\x5f\xd7\x23\x08\x3d\xd4\xc2\xa1\x7c\x7d\xd2\x5e\x08\x30\x6f\xdc\x9f\xe7\xb7\x97\x1b\x98\xf6\x99\xcb\x2d\xb4\xdf\xcd\xb9\xef\x99\x39\x30\x39\x61\x4e\xef\xe1\xbf\xd3\x60\xef\xe3\xc0\xd3\x70\xcb\xbd\x09\x86\x6f\x77\xbf\x39\x8a\xee\x9b\x95\xbf\xaf\x5d\xe5\x55\x30\x44\xdd\xa6\x57\xc1\xea\x2c\xba\x3a\x8b\xfe\xd7\xcf\xa2\x27\x7f\x0b\x52\x1d\x47\x7f\xef\xf3\xc1\x16\x75\xa8\x7f\x02\x00\x00\xff\xff\x38\x50\x9e\x8d\xb8\x26\x00\x00"),
          path: "http-api-test.tml",
          root: "http-api-test.tml",
        },
      
        "http-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\xdb\x3e\x12\x3f\x3b\x40\xbe\xc3\xfc\x75\x28\xa4\x42\x95\xb1\x57\x15\x39\x24\x8d\x9b\x66\xdb\x3a\x41\xe2\xec\x62\xb1\x58\xb4\x8c\x34\xb6\xd9\xc8\xa4\x4a\xd1\x4e\xbd\x86\xbf\xfb\x82\x2f\x3d\x0c\x3f\xd3\xc4\x69\xb3\xed\x21\xaa\xf9\x18\xce\xcc\x6f\xf8\x9b\x11\xa9\x76\x1b\x4e\x48\x72\x87\x2c\x85\x14\xfb\x94\x61\x01\x84\x01\x65\x12\x45\x9f\x24\x08\xf7\x43\x9a\x0c\x81\x64\x19\xbf\x2f\x40\x0e\x11\x3e\xf4\x7a\x97\xc7\x97\xe7\x20\x39\xa4\x74\x82\x42\xea\xd6\x3e\x65\x24\x03\x9e\xa3\x20\x92\x72\x06\xbc\x7f\x78\xd0\x6e\xeb\xae\x01\x9d\x20\x83\x77\x57\x37\xa7\x20\xf0\xfb\x18\x0b\x09\x7d\x2e\x74\xd7\x6c\xf6\x7d\xcc\x25\x42\x74\x2d\xc5\x38\x91\xd1\xc5\xed\x37\x4c\x64\xd4\x25\x23\x9c\xcf\x41\x4e\x73\x8c\xa0\x37\xa4\x05\xd0\x02\x72\xc1\x27\x34\xc5\x14\x6e\xa7\x7a\xee\xb8\x40\x11\x1d\x1e\xa8\x41\xa5\x05\xa5\xda\xb3\xc3\x83\x56\xeb\x14\x33\x94\xe8\x27\x9c\x49\xfc\x21\xa3\x77\xe6\x19\x42\x21\x05\x65\x83\x00\x50\x08\x2e\xd4\xc0\x33\x94\xab\x47\xf9\xb3\x99\x53\xef\x92\x24\x77\x64\x80\xf3\x79\x54\xb5\x35\x54\x0e\x8d\xcc\x40\x09\xbd\xc9\x53\xb2\x7a\xf5\x10\x66\xb3\xc8\x0c\x39\x4e\x94\xc7\x9a\xc2\x1b\x3d\x8d\x25\x9a\x6a\x1f\x67\xd9\xea\x15\xdc\x93\x32\xa9\xff\x04\xe0\xff\xfb\x3f\x3b\x59\xa3\x67\x56\x26\xbd\x13\xb8\xd4\xa4\xd9\x2c\x32\x5d\xcb\x4c\x69\xf4\x2c\x98\xf2\x30\xdf\xce\x0f\x0f\x74\x70\xad\x18\x78\x85\x09\x17\x69\x51\x85\xb3\x0e\x24\x15\xaf\x02\x73\x81\x05\x32\x13\xb2\x02\x8b\x9c\xb3\xc2\x05\xa8\xe4\x40\xea\x01\xaa\x97\x20\x59\x06\xc2\xca\xe3\x7d\x3d\x4d\x0b\xdb\x41\x6f\x17\xa3\x1b\xb4\x2d\x74\x1f\xa8\xc0\x85\x4b\x32\x40\xe5\x7b\xf8\xfa\xad\xe0\x2c\xf6\x72\x32\x40\xef\xab\xea\xb9\xb2\x3a\x5f\xa2\x58\x1c\x24\x9a\x5d\x66\x7c\x8f\x4b\x92\xb9\x25\x6a\x83\xa5\x6a\xff\x62\x2d\x73\xa2\xcd\xa8\x9d\x42\xa4\x5a\xbc\x94\xe4\xd0\x71\x3c\x51\xc1\x60\x4d\x34\x8c\x32\xe4\x59\x6a\x08\x65\x28\x65\x0e\x24\xa7\x30\x24\x2c\xcd\x50\x14\x9a\x1e\xcc\x76\xa7\x6c\xa0\x99\x43\x4b\x2c\xd9\xa5\x28\x09\xa4\x24\x85\x2d\x98\xc4\xe2\xe0\xf4\x32\xda\x68\xa2\x30\x82\xb9\x70\x3c\xa2\xda\x46\x28\x05\x4d\x0a\xb0\xcf\xe8\xb3\x79\x56\xf6\x75\xf1\x1e\x04\xca\xb1\x60\xca\x36\x86\xf7\xa5\x64\xca\x0a\x49\x58\xa2\x28\x4a\xe9\xdf\xd0\xb3\x5c\x8a\xa8\x65\xda\x6d\x2b\x3f\x3a\x3c\xe8\x8f\x59\xa2\x84\xfa\xa3\xc5\x35\x43\x58\x54\x30\x80\xd7\x6e\x31\x6d\x80\xd1\x03\x5e\xd9\x46\xdd\x56\x99\x15\x97\xf3\x43\xd3\x61\xe5\xc7\x30\xd2\x0d\xf3\xca\x28\xb3\x59\x55\xc8\x23\x9d\x98\x54\xa0\xe1\x71\xfb\x42\x72\x48\xcc\x10\x63\xf2\x5a\xb7\x47\x4a\xa4\x16\x7b\xc5\xc7\x12\x63\x68\xcf\xf4\x7f\xe6\xed\x38\x1f\xdf\x66\x34\xf9\x42\x8d\x0f\x3e\xa3\x1c\xf2\x34\x86\xcb\x8b\xeb\x9e\x6e\x38\xb9\x38\xfd\x57\x0c\x7f\xbf\xbe\xe8\x1a\x11\xda\x37\xbe\x8a\x11\x67\x77\x00\x8e\x8c\xe4\x0f\x78\xad\x74\x1c\x4b\x9a\x39\x4a\xb2\x24\xa9\x9d\x33\x82\xf8\xa8\xf4\x68\x17\xef\x7b\x82\x24\xe8\x2b\x3e\x4b\xb1\x8f\x42\x45\x5e\xe4\xba\x3b\x23\x2a\x7d\xf7\xe3\x9c\xf5\xb9\xef\xd9\xf5\x2c\x8d\x79\x41\xe8\xba\xff\x49\xe5\xd0\xc8\x1a\x45\x1d\x96\xfa\x41\x10\x28\x27\xb6\x5a\x89\xfc\x11\x7d\x40\x92\xa2\xf0\x83\xe8\x1a\xa5\xef\x69\xad\x98\x7c\xd3\x9b\xe6\xe8\x85\xe0\x91\x3c\xcf\x68\xa2\x43\xb9\xad\xf6\x8f\x67\x67\x6e\x50\xa5\x04\xc7\x60\x61\x41\x4a\x17\x74\x7a\x4f\x31\x4b\x8b\x72\xaa\xfe\x69\x03\xc2\x1b\x8b\xcc\x8b\x41\x29\x78\x65\x84\xf8\x41\x74\x73\xf5\x49\xa1\x47\xd9\xc0\x0f\x4c\x38\x38\x43\x26\x44\x00\x65\x09\x1f\xa9\x28\xde\x81\x14\xcc\x6c\xda\x57\x20\x28\xe7\x2b\x13\x95\xe7\x4f\x31\xe1\xca\x2b\x6a\xfd\x13\x9e\x4e\xfd\x20\x88\x4c\x9b\xff\xca\xad\x13\xbc\xd5\xb3\xfe\x3a\x02\x46\x33\xb0\x7a\xaf\xf4\x4b\x47\xa1\xdc\xf7\xbd\xf7\x84\x66\x98\xaa\xe8\xcc\x89\x28\x50\xfd\x25\x23\x15\xbd\x29\x8c\x45\x16\xfd\x83\x64\x63\x2c\xb6\xf7\x53\xab\xe5\xe9\xf8\xf1\x62\xa5\x4c\x58\x36\x6e\xe5\xbe\x9a\x03\xcb\x8d\x89\x42\xd8\x8d\xb6\x05\xcc\x2a\xec\x1f\x02\x6e\x4a\x24\xf1\xe2\x12\xb1\xf0\x41\x98\xbb\x2c\x12\x3a\xf0\x94\xb2\x8e\x3e\xa2\x6a\xcf\x85\x50\x21\x56\x43\xfb\xc1\xb8\x25\x25\xf1\x70\xb1\x83\xd1\xcf\x8b\x94\xcb\xc6\x70\x8a\x99\x2a\x85\x77\x80\xcb\x6a\x68\xfe\x6d\xd2\xd3\x2b\x24\x91\xe3\xc2\x8b\x35\x19\x47\xd7\xfa\x97\xc1\x22\x6d\xc2\x57\x6d\x3a\x25\x53\x45\x92\xbf\x64\x4a\x59\xfa\xfc\xec\x6e\x4b\x8d\xdd\xa5\xbc\xfd\x02\xb7\x12\x36\xdb\xc8\x68\x56\x65\x36\x53\x51\x3f\x5b\x66\xbb\xd9\x3e\xb1\xb9\x17\x87\x3d\x25\x36\xb3\xdc\xb3\x26\xb6\x12\x9b\xbd\x24\xb6\xc3\x83\x96\x01\xe9\xfc\x34\x04\x7e\xe7\x36\xcb\x09\x19\xf8\x41\x74\x86\xd2\xce\xf1\x4a\x24\x95\x15\x6a\x67\xfd\xc5\xef\xcc\x1e\xd9\xb8\x43\xba\x1c\xca\xd9\x55\xed\x47\x99\x4d\x4c\xbb\xb2\xc4\xf6\xe4\xed\x36\x03\x17\x3a\x24\x9c\x11\xe7\xa9\x59\x19\x25\x0a\x60\x5c\xbd\xdf\x8c\x99\xb1\x6b\x6e\x5d\xf2\xab\x64\xfb\x9d\xd8\x47\x49\x2b\xc3\xe6\x96\xa7\xd3\xed\x3d\xeb\xd8\x47\xfd\x43\x21\x8c\xf8\x92\x5f\x5a\xad\x1a\xfc\x31\x94\xf1\xf2\x7f\x50\x0a\xac\x33\x7d\x49\x9e\x69\xd4\x07\x15\x75\x85\xd5\x44\xf8\x2d\x6a\xbb\x7d\xc0\xfd\x64\xe5\xc4\x2e\xd5\xc4\x6a\x4b\x97\x96\x19\x5d\x6e\x49\x5e\x0d\x59\x64\x1a\xb5\x54\x39\xc0\x5f\x36\xab\x76\x5c\x63\x8e\xe3\x9e\x2b\x03\x9f\x76\x3e\x75\x7a\x9d\xd5\x59\xd7\x1d\x16\xee\x29\xeb\x9a\xe5\xb6\xc9\xba\x1b\x04\x96\x5e\xdd\xcf\x4b\xe1\xcb\xce\x9c\xfd\x91\xdc\x5e\x13\x97\x3c\x9f\x8c\xc5\x9f\x9c\xae\xab\x98\xaf\xe8\xfa\x11\x5e\x09\x54\x3c\xee\x72\x86\xf8\x58\x2f\x7d\x7f\x28\xfc\xa9\x29\xfc\x0c\xe5\x73\xf1\xf7\x59\xc7\xbc\x41\x5d\x75\xae\x2f\x2f\xba\xd7\x9d\x37\xf5\x57\xa9\x65\x84\xae\x2f\x75\xf6\xc4\xe6\x67\x28\x9f\xf5\x05\xca\xe0\xf2\x27\x03\x3c\xc5\xbb\x93\x56\x63\xdd\xcb\x93\x9a\xab\xe5\x63\xba\xfc\xd4\xcc\x46\x62\x8d\x63\x1f\xe3\xc4\x6c\x80\xf2\xc5\xb2\xec\x86\x73\xac\x8b\x8f\x21\x94\x3e\xff\xd9\x84\xa5\xfc\x58\xa0\xa0\x24\xa3\xff\xd5\x17\x49\xbb\xba\xd4\xdc\x2c\x5a\xde\xbf\x17\x54\xa2\xf8\x3d\xbc\xfc\x3b\xe6\xb2\x8b\x8f\xcb\x92\x58\xe3\xa4\xcf\x5c\x8b\xaf\xcd\x53\x76\x1a\xc9\xb2\x0d\x17\x87\xf6\x5e\x73\x75\xb6\xfa\xe9\x1c\xa5\x6f\xf0\xf7\x97\xa6\x8e\xb3\xec\xb9\x33\x95\x01\x67\x3f\xc9\x6a\x42\x04\x70\x91\xa2\x08\xcd\xe3\x64\x6a\xbf\x8a\xd0\xbd\xb4\x0f\x3c\x5d\x9a\xc7\x7c\x4f\x0f\xf7\x82\xb7\xe0\x12\x97\x1a\x2c\x52\xe1\x86\xf3\x34\xf2\xed\x27\x22\xd5\x98\x96\x9e\x05\x47\x7a\xa4\xd6\x03\x30\x2b\x70\xa1\xd3\x23\x45\xe2\xd9\x6d\x68\xb7\xe2\x46\x4d\x4e\xa6\x0f\xd3\xe5\x64\xba\x56\x1b\xdd\x5d\xdb\x82\x4d\xad\x94\xf7\x14\xb7\xba\x2f\x4e\xd4\xef\x9c\x0c\xb0\xcb\xab\x2b\x84\xda\x97\x08\xce\x12\x91\xe7\x6b\x8b\x83\xc5\x2f\x15\x6a\x4a\x2f\x74\x99\x7c\x7a\xa4\x30\x4b\x38\x9b\x44\xc7\x92\x53\x5f\xe4\xf9\xca\x04\xba\x03\xef\x0b\xd5\x83\x93\xea\x3b\x10\xe1\x4c\x61\xe3\xd1\x2d\x0a\x48\x51\x12\x9a\xed\x70\xb4\xb4\x78\xa6\xd8\x3c\x76\xdb\xee\xaa\xc7\x5e\x6c\x68\x0c\x14\x58\xcb\xbc\x02\x47\xf0\xe6\x6f\xf5\xd0\xc9\x07\xeb\x8b\xb1\x05\x27\x3b\x08\x97\xf8\x36\x1f\x3c\xaa\x6b\xf3\x5f\xcd\x9d\x75\x07\x34\xdd\x58\x2f\xe2\xf4\xd7\x31\x2b\x6b\x39\xcb\xd8\xe1\x02\xb1\x84\xab\xb6\xc6\xa3\x55\x7b\x26\x5d\xed\xbf\xe2\xdb\x57\x4d\xb7\xfe\x0b\x29\xab\xa4\xf2\x68\xec\x5c\x6d\x9a\xec\x80\xb8\xaa\x09\x6d\x47\xfd\xeb\xa7\xd8\xa2\xea\xa6\x34\x20\x8a\x17\x31\xb3\x49\xe4\x45\x96\x96\x2f\xa9\x78\xdc\xbd\x44\xfc\x5f\x00\x00\x00\xff\xff\xaa\x33\x2e\xd7\xf0\x2a\x00\x00"),
          path: "http-api.tml",
          root: "http-api.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}