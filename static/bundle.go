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
        
          "http-api-mock.tml",
        
          "http-api-readme.tml",
        
          "http-api-test.tml",
        
          "http-api.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "http-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x4f\xef\x93\x30\x18\xc7\xcf\x34\xe9\x7b\x78\xe4\x60\x20\x59\xca\x5d\xb3\x83\xf1\xa4\x31\xbf\x19\xe7\x4e\xc6\x64\x1d\x3c\x4c\x26\x94\xa5\x2d\xea\xd2\xf4\xbd\x9b\x02\xc6\x95\xb1\x00\xcb\x2e\x3b\x3c\xf4\xfb\xa7\xcf\xa7\x19\x25\x49\x02\x27\x55\x0b\xc8\x8b\x3f\xba\x91\xa8\x80\x31\x46\xc9\x2f\x2e\x21\xa2\x04\x8c\x61\x5b\x2d\x9b\x54\xb3\xcd\xe1\x84\xa9\x66\x2f\xbc\xc2\xf6\xc7\xda\x8f\xdb\xcd\x0b\xac\x61\x6f\x0c\x54\xfc\xfc\x85\x8b\xac\xae\xda\x59\x2f\x81\xd0\x19\x87\x10\x86\x60\xed\x9e\x92\x09\xbb\xf7\x12\xb9\xc6\xde\xd4\x98\x22\x07\xd6\x8d\xde\xa5\xba\xa8\x05\xfb\xa0\x3a\xa9\xb5\x23\x89\xde\xc9\xb1\x78\x63\xb0\x54\x38\xd0\x7e\xbd\x9c\x71\x44\xef\xc6\xbd\x46\x64\xd6\x4e\x36\xdf\x9d\xb3\x61\xf3\x6e\x34\xa7\xb9\x77\xf2\x81\xe6\x9e\xfe\xa6\x79\xec\xca\x27\x09\x7c\xaa\x79\x76\xb5\x60\x89\xba\x91\x42\x01\x07\x81\xbf\xa1\x10\x4a\x73\x91\x22\xd4\x39\x70\x77\x53\x6f\x1b\x9f\x79\xfa\x93\x1f\xd1\x5a\x36\xfc\xd2\x6d\xc2\x5a\x46\x49\xde\x88\x74\x10\x12\xa5\xb5\xd0\x28\x34\x28\x2d\x0b\x71\x8c\x21\x5a\xee\xbc\x02\x94\xb2\x96\x31\x18\x4a\x02\xf7\x26\xb1\xc4\xea\x81\x86\x6e\x0b\x41\x91\x3b\x37\x78\xb3\x6e\x1f\x3c\xdb\x89\x8a\x4b\xf5\x83\x97\xd1\xb7\xef\x87\x8b\xc6\x7f\x7d\xe3\x15\xbc\x76\x31\xf1\xdb\xf6\xf8\xab\x35\x88\xa2\x6c\x0b\x04\xdd\xda\x1e\xc8\x37\xdd\x4d\x28\x09\xba\x2a\xbd\x91\x8b\x59\x39\x7b\x4a\xec\x15\xa8\xab\xf7\x34\x01\xca\x83\xef\xd5\xf0\xbe\x8c\x81\xfa\x1f\x32\x0a\x6a\xa1\xf3\x5d\x50\x0b\x7d\x9e\x0e\x6a\x61\xfe\x32\x50\x93\xff\x8e\x13\xf8\x7a\xb1\x57\xeb\xd6\xd0\xe7\x36\x95\x39\x4a\x73\x6e\xd0\x5d\x8c\x73\x0d\x9e\xce\x6f\x6e\xf0\x2c\x70\x7f\x03\x00\x00\xff\xff\x61\x7f\x28\xae\xee\x06\x00\x00"),
          path: "http-api-json.tml",
          root: "http-api-json.tml",
        },
      
        "http-api-mock.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x97\xdf\x6f\xdb\x36\x10\xc7\xdf\x0d\xf8\x7f\xb8\x04\xd8\x20\x15\x86\xfc\x9e\x22\x0f\x59\x53\x0c\x01\xd6\x24\x68\xd6\xa7\x61\x18\x68\xf1\x64\x71\x91\x48\x95\xa4\xe4\x04\x82\xfe\xf7\x81\xbf\x64\x59\x72\x9c\x2a\x2d\x30\xf4\x25\x8e\x45\xf2\x7e\x7c\xef\x73\x47\x6b\xbd\x86\x4f\x22\x7d\xbc\xba\xbf\xb9\xab\x50\x12\x2d\x24\x50\xcc\x18\x47\x05\x84\x83\xd2\xb2\x4e\x75\x2d\x11\x76\x39\x4b\x73\x60\x65\x55\x60\x89\x5c\x2b\xd0\x39\xc2\xf0\x54\x26\x24\x54\x52\x34\x8c\x32\xbe\x85\xe5\x62\xbd\x06\x02\xa5\x48\x1f\xa1\x56\x64\x8b\xc0\x38\x68\x54\x5a\xd9\x8d\xb5\x42\xd8\x31\x9d\x5b\x2b\x6d\xfb\xb5\x16\x1a\x21\x79\xb0\xde\x92\xbb\xcd\xbf\x98\xea\xe4\x96\x94\xd8\x75\x90\x6b\x5d\x19\x47\xc9\x72\xa1\x9f\x2b\x9c\x44\xeb\x42\x6c\x97\x0b\x00\x00\xa5\x85\x44\x78\xa7\x9e\x79\x5a\x92\x2a\xf9\x44\xaa\xe5\xa2\x5b\x2e\x6c\x34\xb7\xb8\x1b\x9f\x95\xa8\x6b\xc9\x15\x10\xe0\xb8\x03\xc6\x95\x26\x3c\x45\x10\x19\x90\xb1\x9f\x64\xb9\xc8\x6a\x9e\x1e\xb1\x12\xc5\xf0\x6e\x6c\xd8\x87\xe3\xec\xc3\xaf\xa3\x65\xbf\xda\x07\x7c\x61\xdc\x47\x83\xa0\xe3\x95\xdb\xd1\xed\xa3\xbf\xc6\x02\x35\x7a\x85\xd1\xc9\x2f\xac\x3d\x26\x38\x68\x01\x12\x4b\xd1\x20\x10\xd8\xb2\xc6\x54\x40\x62\x2a\x24\x05\x46\x91\x6b\x96\x31\xa4\xb0\x79\x86\x9b\xeb\x90\x47\x54\x4e\xa2\x8e\xbd\x93\x28\xd5\x4f\x90\x0a\xae\xf1\x49\x27\x1f\xdc\xe7\x0a\xaa\x7a\x53\xb0\xf4\xe6\xda\x08\xce\xf8\x36\x06\x94\x72\x9f\x29\xcb\xe0\x9f\x15\xe4\x44\xc1\xc5\x25\x94\x89\x4d\x2b\xf9\x43\x10\x1a\x85\x73\xf1\x7b\x38\x33\xeb\x83\xe4\xbd\x3c\x59\xa9\x93\x8f\xc6\x58\x16\x9d\x7f\x76\x51\x53\x81\x0a\xb8\xd0\x80\x4f\xcc\x30\x63\x59\x61\x14\x7e\xf9\x7a\xbe\x8f\x24\xee\x45\x72\xff\x04\xb7\x3e\x8b\xd1\x36\xef\x8c\xb3\x62\x2f\xea\xef\xa8\xaf\x8a\x62\x80\x81\x2a\x98\xaf\x7f\x51\x00\x69\x08\x2b\xc8\xa6\xc0\xa0\xa5\xc8\xc0\x22\xd8\xb6\x01\xd5\x7b\x92\x3e\x92\x2d\x76\x5d\xb2\x7f\x76\x80\xef\x49\xb9\x9d\xfb\xe3\x72\x0b\x49\x51\x7a\xad\xfd\xb7\xdf\x9e\xfb\xef\x95\xeb\x28\xbd\x02\x89\xaa\x12\x5c\xa1\xbc\x47\x79\xef\x9f\xc6\x10\xfd\xf5\xf7\x8c\x20\x57\xce\x94\x2d\x68\x1c\x0a\xd4\x10\xe9\xf3\x56\x30\xcb\xda\xb8\x1c\x9f\x09\xdf\x62\x64\x54\x88\x1e\x57\xd0\x18\x5f\x28\x33\x92\x62\xdb\xc5\xb0\x11\xa2\x18\x22\xc1\x32\xc0\x02\xcb\x15\x88\x47\x43\x52\x93\x44\x33\x3c\xc7\xef\xcd\xb1\x81\x35\x57\x77\x97\xc3\x25\x90\xaa\x42\x4e\x23\xff\x60\x65\x1d\xc5\xfb\xcd\x7d\xe0\x03\x5a\xb4\xac\xd1\x53\x16\x87\x65\xbf\xd4\x9b\x29\x90\x07\x9b\xf1\x6a\xc2\x97\xd9\x2e\x19\x36\x66\x9e\x06\x8e\x36\x44\x21\x05\xd3\xb6\x79\xdf\xd1\xb4\xc7\xfa\x35\x64\xbe\xb1\x3d\xe7\x08\x17\x6a\xef\xa5\x73\x15\xc8\x44\xcd\xe9\xcb\xed\xdc\xf7\xfd\x99\xdd\x38\x6d\xeb\x19\xfe\xdb\x6e\xf5\xc3\xa6\x80\xfc\x38\xe0\xc7\x64\x32\x0f\xa1\x7d\x5a\x87\x28\xbd\x2d\x29\x2b\xab\x4a\x6e\x71\x37\x4d\xaa\x24\x3a\xcd\xed\x40\x39\x9f\xe4\xe0\x11\x73\xa9\x1c\x20\xf5\xa5\xa2\x44\x23\xd4\xf6\x43\x4d\xc6\x7d\x7f\xa9\x6e\x59\x83\xdc\x5e\x6b\x0d\x29\x6a\x3c\x09\x95\xb3\xf9\x4d\x5c\xb9\xae\x31\x32\xb8\x43\x57\xa9\xb9\x7f\x0e\xc5\x38\x58\x71\x92\x74\x9d\xbf\x2f\xda\xe5\xa2\x6d\x8d\xbc\x7e\xd7\x8d\x7a\x60\x25\x2b\x88\x04\x33\x38\x86\x63\xe3\xc1\xfc\xed\x71\x3b\xe8\xd6\xe1\x28\x6f\x5b\xc0\x42\x61\x7f\x9c\x8b\xdd\x10\x80\xe3\xec\xce\x43\x02\x5e\x61\x62\x58\xe4\x5b\xa1\x5d\xe3\xac\x19\x6f\x48\xc1\xe8\xb8\xc0\xe6\xc3\x47\x8a\x3e\x52\x94\xd2\x84\xea\x4a\xfa\x42\x10\xa6\x38\xab\x7d\x72\x03\x31\xcc\xc4\x94\x12\xce\x2e\x8d\x1e\xc7\xa3\x7b\xe1\x9e\x74\x0a\xfb\x30\x92\xfb\x5e\x69\xff\x64\x3c\xf0\x7a\xb9\x39\xb5\x6a\x07\x22\x3f\x48\x34\x44\x12\x4a\xc3\x0f\xa9\xf0\xd3\x83\x6b\x11\x58\x1c\x20\xea\x9c\x9f\xe2\xd1\x59\x3c\xce\x63\xc0\xcf\xed\x39\x86\xdf\xc1\xca\x1e\xbf\x37\xce\x43\x4f\xab\x33\xfa\x1a\xad\x78\x28\xe4\x14\x59\xdc\x37\xf4\x04\xdc\x11\x0e\xa9\xf5\x78\x12\x87\x19\x14\xcc\x9f\x5b\x3f\x08\x9a\x90\xd5\x51\x7a\x9c\xb8\x91\x99\x86\xe3\x79\x10\xdb\x5d\xeb\xf5\xe9\xb6\x18\xcc\x41\x7c\x62\xda\x50\x36\x23\xd3\x17\x86\xa6\xb3\x09\x4c\x63\x19\x30\xfd\xf2\x5a\x6f\x4e\x41\xf5\xb6\x67\x51\xe7\x3c\xcb\x37\x8d\xd7\xef\xb8\xef\xdd\xdf\xf5\x1a\xfe\xbc\xbb\xbe\x8b\x28\x36\x58\x98\x57\x8b\xf8\xa2\x5f\xb8\x6b\x50\x4a\x46\x11\x8c\x1a\xf6\x8d\xc3\xe6\xcb\xc3\xa5\xbc\xcb\x89\x06\x95\x8b\xba\xa0\x90\xdb\xdf\x59\xc9\x08\x84\xef\xb8\x38\xcd\x4c\xbd\x09\x2f\x9e\x48\xcd\x30\xed\x0c\x3b\xc8\x69\x37\xc1\x68\xd4\xa8\x3d\x46\x27\xdb\xc9\xaf\x86\xf9\xf5\x06\x84\x32\x29\xca\x01\x42\xce\x9e\x9d\xfd\x01\xa1\x57\xfb\xf9\x7f\x9e\x75\x66\x68\xfc\xcc\x20\xfc\x17\x00\x00\xff\xff\xa1\xf0\xff\x07\xc9\x10\x00\x00"),
          path: "http-api-mock.tml",
          root: "http-api-mock.tml",
        },
      
        "http-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x4d\x6f\xe3\x36\x13\xbe\x1b\xf0\x7f\x98\x37\xbe\x24\x8b\xac\x9c\x7d\xd1\x5e\x0c\xf4\x90\xcd\x66\x53\x17\x6d\x62\xc4\xde\x53\x10\x54\x34\x39\xb6\x98\xa5\x48\x95\x1c\x39\x31\x04\xfd\xf7\x82\x94\xec\xd8\x5e\x7f\xa6\xee\xa2\x28\x9a\x43\x08\x53\xf3\xfd\x3c\x33\x1a\x15\x45\xd4\x27\x9b\x73\x8a\xee\x86\x4f\xc8\x29\xba\x65\x29\x96\x25\xfc\x3c\x18\xf4\xe0\xb2\xd7\x85\x66\xe3\xa7\xed\x7f\xcd\x46\xb3\xf1\xf0\xbf\x87\x1b\x03\xf7\x98\x19\x4b\x70\xc5\xac\x78\x3c\x4d\x88\x32\xd7\x69\xb7\xc7\xc6\x86\x6b\xce\xac\x88\xb8\x49\xdb\x43\x26\xc6\xd8\x2e\x8a\xa8\xc7\xf8\x57\x36\xc6\x1e\xa3\xa4\x2c\xcf\xb6\x68\x54\x3f\xbf\x55\xf1\x9e\x77\xc6\x2f\x1d\x30\x60\x39\x19\x18\xa3\x46\xcb\x08\x05\x78\x4f\xc0\x32\x09\x23\x63\x81\x12\x04\x17\x4c\x40\xbc\xd6\x5a\x6d\x32\x8e\xbc\xbf\x41\x82\xc1\x2c\xbe\x64\xc8\xc9\x05\xed\xdc\xa1\x05\x32\x20\xd3\x4c\x61\x8a\x9a\x80\x69\x01\x99\x35\x13\x29\x30\x48\x0c\x19\xff\x8a\x5a\x80\xd4\x84\x76\xc4\x38\x7a\xf1\x5a\x40\x40\xae\x05\x5a\x25\x35\x82\x18\x56\x22\x8c\x93\x34\xba\xe3\x1d\xc6\x71\x3c\x36\xcd\x06\x4d\x33\x84\x8f\xab\x66\x8a\x66\x03\x00\xe0\x13\x2a\x24\x3c\xe5\x46\x13\xbe\x50\x74\x55\x9d\xe7\x3e\x2d\xa9\xc7\x67\x80\xd6\x1a\x5b\x89\xde\x20\x6d\x96\x3b\x7d\xcd\xbf\xae\x74\x59\x46\x1b\x2a\x7c\x5e\x59\x3d\xab\xcc\x7e\xc9\x04\xdb\x1c\xc1\x39\x14\x45\x54\x89\x5c\x86\xcc\x96\xcd\x2f\x3d\xa9\x9c\x94\xe5\x6a\xd4\x97\x4a\x6d\x36\x3f\x3b\xa5\xa6\xf0\xef\x0c\x4e\x1f\x1e\x0f\x4a\x26\x68\x2e\x66\x74\x65\x71\x6d\x46\x45\x11\x55\x8f\xd6\x65\xb2\xf4\xe4\x35\x93\xb7\xd5\xb5\x0c\xe0\x7b\x0e\x74\xc9\xf3\xcd\x38\xac\xf8\x36\x32\x4a\x99\x67\xa9\xc7\x90\x22\x25\x46\xb8\xc0\x63\x64\x3c\x01\xd4\x22\x33\x52\x53\xa0\x4e\xab\x55\x67\x01\x05\xf4\xee\xfa\x03\x68\x6f\xf0\xd9\x86\xd2\x4b\xb7\xe0\xb7\x60\xaf\x03\xf1\x28\xd7\x1c\x4e\x7d\x8b\xbc\xf3\x9d\x74\xd9\xeb\x9e\xcd\x2b\x42\x2f\xf0\xce\x77\x50\x4e\x52\xcd\xea\x52\xa3\x15\xa2\xad\x7d\x5a\xe4\x28\x27\x75\xc8\x73\xae\x5b\xe4\xc6\x0a\x30\xa3\x70\xbd\xa9\x64\x10\xe8\xfe\x9c\x48\x9e\xf8\x06\x16\xa8\x24\x4e\x50\x04\x9d\x66\xe3\x97\xfe\xdd\x2d\x04\x60\x34\xf9\x46\xf2\xb7\xb3\x7e\x8f\x60\x90\x48\x07\xcf\x52\x29\x90\x1a\x28\xb7\x1a\x2c\x86\x83\x81\x45\xe7\xbb\x56\x4e\x7c\xc7\x33\xca\x1d\x70\x23\x30\xf4\xf5\x7b\xb8\x0e\x1d\x8d\x02\xee\xf1\x8f\x1c\x1d\xc1\x47\x23\xa6\x75\x07\xfa\x74\x6b\x5a\x54\x6e\xdf\x0f\xa6\x19\x76\x80\x65\x99\x92\x9c\xf9\xd8\xdb\x4f\xce\xe8\x39\x62\x71\x1c\x57\xbf\x8b\x42\x8e\x60\x39\xc9\xae\xab\x30\x28\xcb\xa2\x80\x94\x65\xf7\x4c\x0b\x93\x86\xa4\x96\x05\x2b\x31\x38\x19\x3a\xa3\x4f\xe0\xe4\x29\x1c\x5e\x0b\x95\x43\x7f\x1e\xca\xc4\xa2\x40\x2d\xca\x57\x5e\x2d\x24\xdd\xaf\xea\x71\x65\x04\xd6\xf1\x37\x1b\x9f\x99\x54\xb9\xc5\x0e\xfc\x78\x71\xd1\x6c\xf4\x73\xce\xd1\xb9\x0e\xfc\xff\xe2\xc3\x3a\x0b\xf7\xe8\x32\xa3\x1d\x1e\xab\x6e\xab\xa5\xd9\x50\x8d\xb9\x6a\xab\x05\xdd\xdb\xcf\x77\x9b\x49\xbe\x07\xc5\xbb\x7a\x64\x76\x13\xdc\x4b\xd5\xa4\x72\x40\x86\x98\xf2\x7c\xae\x98\xed\x80\x4d\x98\x54\x6c\xa8\xd0\xf3\x6f\xfe\x92\xf1\x7c\xde\x7b\x0a\xd4\x11\xaf\x12\xf3\x20\x8c\x2e\xbe\x03\x46\x95\xda\x49\x28\xc1\x49\x07\x3e\x5c\x9c\x2f\x4e\xad\x56\x0b\x6e\xae\xb7\x0c\x9d\x4e\x96\x0f\x95\xe4\xbf\x4b\xb1\x07\x32\xe1\xbd\xb5\x0b\x98\x1b\x24\x8f\x8b\x0d\x63\x87\xc1\x58\x4e\xfc\x90\xfc\x66\xe4\xac\x5d\x18\x02\x42\x23\x6b\xd2\xa5\x71\x52\xc3\x5c\x9b\xf1\xf3\xcc\x8f\x30\x97\x2b\x02\xe6\x5d\x78\x6a\x36\x1b\xb6\x2e\x6b\x04\x5d\xf2\xdb\xc0\xca\xcc\x8b\x5f\x33\x8d\x21\x63\x96\xa5\x48\x68\xbd\x81\x20\xe6\x2f\xa8\xde\x20\x04\x6a\x92\xa3\x69\x78\x50\xc5\xbd\x69\x3a\xf5\x66\x76\xdc\x9c\x0b\x1e\x8c\xc5\xa2\xfe\xc5\x2e\xff\x1e\x0c\x7a\x4b\x97\x6f\x25\xd5\x7e\x54\x0a\xcb\xc4\x61\x6c\x52\x6a\xde\xe1\x6f\x62\xd2\x3f\xbd\x99\x1f\xf6\xc6\xe2\x7c\x6f\xc9\xc7\x45\xd8\x7a\x5f\x8e\x35\x0b\x66\xcb\xe6\x2e\x00\x2b\x39\x60\x44\x98\x66\x7e\x57\x37\x90\xd7\x57\xff\x8e\xd9\x00\xcf\x92\x92\x65\x7b\x01\x0f\x5b\xcf\x88\xa1\x11\xd3\xa3\xce\x8f\xbf\x61\x35\x5a\x5a\xfe\xb7\xad\x46\x4b\x82\xbb\x57\xa3\x03\x3f\x37\x8e\xb5\x1a\xfd\xb0\x48\xf9\x4f\xd7\xbf\x5e\x0f\xae\x8f\xc4\xfa\xd9\x47\xde\x3e\x63\x4b\x04\xd9\xff\xde\x81\x47\x84\xf3\xcf\x00\x00\x00\xff\xff\xcd\x02\x02\x93\x29\x11\x00\x00"),
          path: "http-api-readme.tml",
          root: "http-api-readme.tml",
        },
      
        "http-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\x4a\x2b\xcd\x4b\x56\x08\x49\x2d\x2e\xa9\xae\xd6\x0b\x2e\x29\x2a\x4d\x2e\xd1\xf3\x4f\xca\x4a\x4d\x2e\xd1\xf3\x4b\xcc\x4d\x05\x13\xb5\xb5\x8e\x01\x9e\x1a\x25\x0a\x5a\x25\xa9\xc5\x25\x99\x79\xe9\x7a\x21\x9a\xd5\x5c\x9c\xfa\xfa\x0a\x79\xf9\x25\x0a\x95\xa9\x25\x0a\x99\xb9\x05\x39\xa9\xb9\xa9\x79\x25\xa9\x29\x7a\x5c\xb5\x5c\x20\x00\x08\x00\x00\xff\xff\xea\xe9\x09\x4d\x58\x00\x00\x00"),
          path: "http-api-test.tml",
          root: "http-api-test.tml",
        },
      
        "http-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xe3\x46\xd2\x7d\xa6\x01\xff\x87\x0a\x1f\x02\x2a\xd0\xd0\xf8\x5e\x19\x18\xf8\xec\xb1\xc7\xf1\x26\xb1\x0d\xdb\xb3\x8b\xc5\x62\x91\xb4\xc9\x92\xdc\x19\xaa\x9b\x69\xb6\x7c\x59\x41\xff\x7d\x51\x7d\xe1\x45\x43\xc9\x94\x6c\xcb\x33\xb3\xf1\xc3\x68\xc8\xbe\x55\xd5\xa9\x3a\xa7\x49\xf6\xde\x1e\xcc\x66\xf1\x95\x56\xd3\x54\xc7\xe7\x37\x7f\x60\xaa\xe3\x33\x36\x41\xf3\xcf\x7c\x7e\xc8\xd2\x4f\x28\x32\xc8\x70\xc4\x05\x96\xc0\x04\x70\xa1\x51\x8d\x58\x8a\x70\x7f\xcb\xd3\x5b\x60\x79\x2e\xef\x4b\xd0\xb7\x08\x3f\x5d\x5f\x5f\x1c\x5c\x9c\x82\x96\x90\xf1\x3b\x54\xda\xdc\x1d\x71\xc1\x72\x90\x05\x2a\xa6\xb9\x14\x20\x47\xbb\x3b\x7b\x7b\xa6\x69\xcc\xef\x50\xc0\xfb\xcb\x8f\x47\xa0\xf0\xcf\x29\x96\x1a\x46\x52\x99\xa6\xd9\xec\xcf\xa9\xd4\x08\xcb\x6d\x03\xfd\x58\x60\x0c\xd7\xb7\xbc\x04\x5e\x42\xa1\xe4\x1d\xcf\x30\x83\x9b\x47\x33\xc1\xb4\x44\x15\x9b\x95\xfe\x9f\x4f\x8a\x1c\x27\x28\xf4\xee\x0e\x8d\xe9\xe5\x71\xe5\xe6\x6c\x77\x27\x08\xde\xcb\xa9\xd0\x51\x2a\x85\xc6\x07\x1d\xbf\xb7\xbf\x03\x88\xb8\xd0\x43\x40\xa5\xa4\x1a\x50\xb7\x23\xcc\x51\xe3\x62\xbf\x21\x94\x5a\x71\x31\x1e\xd8\x9e\xd4\xf1\x04\x3f\x9b\xad\xee\x15\xd5\xf6\x5d\xb0\xf4\x13\x1b\xe3\x7c\x1e\xaf\xb2\xb9\x69\xc2\xc7\x22\x63\xcb\x4d\x18\x92\xef\xb6\xcb\x41\x4a\x68\xb4\x57\x68\xb5\xd8\x75\xe6\xf3\xb6\xd9\x07\x79\xbe\x7c\x72\xff\x6b\xc2\xc2\x05\x45\xe8\x5f\xff\x5e\xdf\x9b\x85\xa8\xbe\x57\xd8\xe9\xd2\x6c\x16\xdb\xa6\x2e\x57\x5a\x2d\xb5\x2b\xcf\x88\xed\x7c\x77\xc7\xa4\xd3\xaa\xde\x54\x01\xab\x8a\x05\x1f\x0a\x59\xa2\xc9\xcf\x09\xea\x5b\x99\x7d\x9e\xb7\xb7\x5a\x17\x70\x63\xb3\x30\xee\x91\xb0\x66\xc9\x56\xb6\x52\x72\xfd\x40\xd3\x4c\x35\xcf\xeb\x64\xf5\x28\x9e\x8a\x91\x5c\xd5\xee\xc2\xbd\xa2\x87\xcb\xb1\x15\x3d\x5c\x21\xac\xe8\xe1\x52\x69\x79\x8f\x3e\xe1\x26\x57\x80\xf2\x82\x71\x51\x52\x44\xd9\xbb\x8c\x69\x06\x0a\xc7\x4c\x65\x5c\x8c\x41\x61\x2a\x55\x56\x02\x17\x90\xdd\x54\xd4\xd3\x8e\x69\xbf\x44\xe8\x83\x85\xb1\xa7\x34\xad\x40\x48\xc0\xb5\xd4\x2c\x27\x74\xe0\xf7\x3f\x4a\x29\x92\x50\xd3\x8d\xf0\xf7\x7e\xee\x5d\x3a\xe3\xab\x84\xb2\x96\x6b\x09\x0a\x0b\x85\x25\x0a\xcb\xb0\x0a\xcb\x42\x8a\xd2\xf3\xa9\x96\xc0\x9a\x7c\x6a\xd6\x61\x79\x5e\x05\x43\x8e\xcc\xb0\x57\x0b\x83\xb7\xbb\x19\x89\x0b\x36\xc6\x66\x20\x0a\x36\x46\x8a\x43\x00\x97\xce\xfa\x0b\x54\x8b\x9d\x54\xbb\xc9\xf6\x37\x31\xbd\xac\x70\x6d\x87\xf6\x37\xe7\xa3\x9f\xda\xf6\x5a\x9f\x84\x6a\x0b\xaa\xe9\x3c\x62\x5e\xe5\x6a\x54\x9c\x9f\xb6\xc4\x6f\x65\x9e\x95\x75\x25\xb3\x82\xc3\x2d\x13\x59\x8e\xaa\x34\xe2\x66\xeb\x9d\x72\x93\x74\xcf\xcc\x58\x69\x63\x59\xc9\x5f\xc5\x0a\x7d\x75\xd0\xc1\xe2\x8d\xb3\x26\x19\xed\x9a\xa0\x56\x3c\x35\xe5\x41\xbf\xf1\xaf\xf6\x97\x9a\xec\xc2\x52\xf5\x51\xc4\x3a\x00\x67\x78\x0f\x0a\xf5\x54\x09\x72\x5e\xe0\x7d\xb5\x2a\x17\xa5\x66\x22\x25\xf1\x25\x07\x5b\x8e\x54\x6b\x31\x9a\x6b\x6f\xcf\xd9\x13\xef\xee\x8c\xa6\x22\xa5\x49\xa3\xc9\xa2\x8d\x43\x4f\x84\x7d\x0c\x1c\xc0\x0f\xde\x0e\xe3\xb7\x35\x11\xbe\x77\x37\xcd\xbd\x2a\x1a\x09\x4c\x86\xf6\x86\xb7\x2b\xf1\x6b\x99\xfb\xf3\xda\x5d\x53\xd4\x0a\x53\xe4\x77\x96\xd4\x0d\xae\xbe\xbe\xb4\x84\x31\x6a\x57\x5a\xc0\xa9\x2f\x61\x48\xe5\xc6\xee\x18\xcf\xd9\x4d\x8e\xad\xc2\xb3\xc5\xf3\x14\xa8\x66\xe7\x62\x96\xbf\x94\x53\x8d\x09\xec\xcd\xcc\x7f\xe6\x7b\xb4\x84\x69\xf8\xd5\x08\x48\x02\xa7\x67\x1f\xce\x6d\xcf\xe3\xab\x8b\xf3\xb3\xab\xe3\x77\x87\xe7\x47\xff\x4c\xe0\x6f\x57\xe7\x67\x2e\xb8\x11\x65\xa1\x8f\xce\xc0\x78\x14\xa5\xfa\x01\x96\xb1\xaf\x09\xe0\x04\x92\xfd\x0a\x90\x33\xbc\xbf\x56\x2c\xc5\x88\x14\x39\xc3\x11\x2a\xca\xec\xd8\x37\x1f\x4f\xb8\x8e\xfc\x85\x99\x3e\x74\xab\x99\xab\x70\x30\xac\x66\xfa\x07\xd7\xb7\x76\xaa\x49\x7c\x2c\xb2\x68\x30\x18\x50\xa0\x01\x00\x52\xfd\x10\xff\x84\x2c\x43\x15\x0d\xe2\x2b\xd4\x51\x68\xcc\x12\xfa\xdd\xf5\x63\x81\xe1\x10\x42\x56\x14\x39\x4f\x4d\xad\xec\x51\x81\x86\xd5\xd8\x27\xac\x71\x20\x5a\xcc\x1c\x98\x59\x38\x18\x36\xad\xfa\xc0\x31\xcf\xca\x6a\xa0\xb9\x9c\xd9\xd9\xfd\x5f\x38\x55\x79\x98\x18\x43\x2f\xed\x64\xd1\x20\xfe\x78\xf9\x0b\x01\xc9\xc5\x38\x1a\x0c\x6d\xff\xb9\x75\xca\x5e\x18\x62\x32\xfb\x08\x8a\x28\x19\xea\x73\x2e\x76\xfb\x4a\xfd\xe0\x01\x88\x68\x20\x0d\xe2\x23\x33\xe0\xbb\x7d\x10\x3c\x87\x86\x19\x4b\x1d\x3d\x26\xe4\x46\x51\xf8\x81\xf1\x1c\x33\x9f\x98\x6b\x13\x9f\xcb\xe4\x94\x2c\x5b\x3f\x40\x26\x48\x26\x87\xc2\x84\x1c\x18\x76\x34\xf7\x8c\x61\x23\x8e\xfe\xd2\x95\x34\x2a\xe5\xc2\xec\x1b\x5d\xb4\x92\x7d\x33\x2d\x65\x7e\x44\xa9\x1d\x5f\x69\xa6\xa7\xe5\xf9\xcf\xc3\x27\x65\x7b\x66\x64\x25\x31\x58\xcd\x07\x3f\xbe\x4c\xf0\x4b\x54\x9c\xe5\xfc\x3f\x98\x6d\x8c\x83\x11\x7b\x27\xef\xf7\x8a\x6b\x54\x8b\xc5\xf4\x65\xa2\x12\x3c\x51\x90\x5e\xf5\xe1\x08\x73\x7a\x56\xec\x5b\x8d\x41\xe0\x2c\xb5\x7f\x2b\xed\x0d\x82\xb0\x34\x09\x10\x26\xd0\x4e\x87\xdd\x9d\xa0\xe1\x84\x73\x40\xf0\xbc\x66\x7d\xbb\x0b\x5e\xc9\xfb\xa9\xed\x62\x25\xf0\x39\x94\x9e\x14\xd3\x9b\x9c\xa7\xbf\xf1\xac\x45\xec\x17\xe7\x57\xd7\xe6\x46\x93\xcf\x69\x8a\x2e\x4e\xf7\x0f\x49\x5b\x62\x75\xbb\xdc\x02\x64\x1d\xb4\x1e\x04\x9b\x92\x7a\xf0\x54\x06\x55\x08\x6d\x42\xea\xa4\xfc\x7d\x52\x3e\xf0\x79\x12\x04\x77\x4c\x01\x17\xa9\x9c\xd0\xd6\x66\xfd\x67\x4f\x3b\x49\x4d\x54\xe4\x29\x01\x70\x84\xa9\xa4\xe0\x90\x19\x87\x32\x7b\x8c\x06\x83\xd8\xde\x8b\xbe\xf7\xcb\x7d\xc6\x47\x64\xfe\x1a\x5c\x54\x30\x55\x22\xfd\xcb\x26\x94\xc9\x19\x4c\x55\x1e\xff\x9d\xe5\x53\x2c\xfb\x87\x2b\x08\x16\x28\xc4\xdd\xec\x15\xc5\x46\x1c\xab\x9d\x99\x21\x0c\xb3\xd5\xea\x81\x36\x65\xff\x26\x18\xd3\xf3\x60\x98\x54\xc0\x0d\x37\x82\xde\x33\xf0\x12\x11\xaf\x4a\xaf\x56\xf1\x21\xd4\xd8\x35\x70\xdf\x18\xc1\xb4\xa2\x23\xa9\xd6\x70\xff\x6d\x31\xdb\x9c\xe3\xd7\x22\xf9\x6e\x96\xb7\xa8\x64\x6d\x20\x57\xef\x13\xfc\x90\x4a\x71\x9f\x5b\x77\x99\xf5\xbb\x9a\x6f\xbb\xc0\x2d\x85\xad\x4b\xef\xec\x3b\x9d\xb7\xd5\xbb\x8f\xfd\xe5\xce\xbf\xe6\xdc\x92\xdc\xd9\xe5\xde\x54\xee\x2a\x80\xb6\x22\x77\xbb\x3b\x81\x05\xe9\xf4\x68\x08\xf2\x93\xaf\x98\x43\x36\x8e\x06\xf1\x09\x6a\x37\x26\xac\x90\x24\x2f\xa8\xbc\xbe\x93\x9f\x6c\xa1\x3c\x59\x26\x67\x12\xaa\xd1\xf5\x6b\x02\x2e\x9c\x4e\xad\x4b\x15\xfd\xb9\xdc\x57\x84\x54\x26\x25\xbc\x13\xa7\x99\x5d\x19\x35\x2a\x10\x52\xc3\x48\x4e\x85\xf5\x6b\xee\x42\xf2\xd9\x1e\x60\xcd\x57\xe9\x2f\xbc\x07\x58\x8b\x89\x68\xb6\x2a\x7b\x6e\x64\xf6\xd8\x3f\xc0\x9e\x89\xe8\x0f\x95\xb2\xd3\x57\x5c\x13\x04\x8d\x2c\x48\xa0\x4a\x9b\xff\x81\x0d\xc2\x2a\xd7\x3b\x34\xa7\xb5\x6b\xa8\x19\xac\xb9\x6b\xa8\x8b\xee\xab\xd8\xfb\x6d\x03\xf8\x2f\xe3\x41\x72\xa9\xa7\x9d\x9b\x8f\x33\xe9\x58\xbf\xfb\x49\x93\x16\xab\xba\x44\x5d\xe3\x1a\xdf\x9b\xec\xb7\x94\x37\x55\xe6\xa3\xe3\x5f\x8e\xaf\x8f\x97\xab\xb1\xff\xee\xb9\x25\x35\xb6\xcb\xf5\x51\xe3\x27\x26\xac\x42\xbb\x9d\x47\xc8\x6f\x5b\x51\x47\x13\xdd\xdf\x12\x2f\xaa\xaf\x46\xeb\xaf\xce\xdf\x75\xce\x77\xf1\xf7\x0b\x3c\x39\x50\x66\x6e\xf8\xd6\xf0\x2f\x62\xdf\x26\xb1\x57\xde\xad\x45\xeb\x27\xa8\xdf\x94\xd3\x4f\x8e\xaf\xd7\xfb\x6a\x64\xce\xac\x6c\x89\xe1\x4f\x70\xf1\xdb\xc3\x76\x1f\xb6\x2c\x38\x7f\xa9\xc2\x6b\x3c\x67\x19\x33\x56\x3d\x68\xd1\x58\x33\x3f\x66\xdd\x2f\xdc\x5c\x26\x76\xf2\xee\x4b\xbc\x6c\x7b\xc6\x77\xb3\xaf\x83\x79\x9f\x78\x0f\x76\xfe\xf3\x10\x2a\x08\x9e\xab\x64\xaf\xf9\x1d\xec\x2b\x08\xf5\xd7\x28\x72\xad\x0f\x64\xdd\xaf\x0b\xed\xf1\xad\x95\x02\xe6\x86\xb1\x3c\xef\x73\x88\xc5\x1d\x91\x58\x2e\x63\xcf\x16\x2f\x73\x72\x71\x7b\xfa\x75\x90\xe7\x6f\x2d\x61\x16\xa1\xed\xa8\xd8\x1d\x53\x20\x55\x86\x6a\x68\x7f\x0e\x1f\xdd\x69\x50\xd3\xca\x47\x20\xb3\x4e\x81\x8b\x42\xd3\x3d\x1c\xfc\x08\x5e\xd1\xa8\xb3\xca\x94\xef\x2e\xb3\x38\x72\x47\x63\xeb\x3e\x81\x19\x05\xfb\xa6\xa7\xb1\x03\x30\x2f\x71\xa1\x31\x64\x65\x1a\xba\x5a\x74\xf5\xf8\xa4\x25\x87\x8f\x9b\xd9\x72\xf8\xb8\xd2\x1a\xd3\xdc\xa8\xc3\xb6\x55\x14\x3d\x62\x59\x7f\x3a\x92\xae\x0b\x36\xc6\x33\x59\x7f\x8c\x68\x9c\x8f\xf3\x9e\xa8\xa2\x58\xb9\x6b\x58\x3c\x3f\xd7\x30\x7a\xa1\xc9\x0a\xed\x3e\x61\x96\x4a\x71\x17\x1f\x68\xc9\x23\x55\x14\x4b\xf5\x74\x0d\x05\x50\xd4\x82\x77\xf5\x39\x45\xe5\x5d\x11\xd3\xc9\x0d\x2a\xc8\x50\x33\x9e\xaf\xf1\x3a\x6a\xf1\x8d\x64\xfb\xa5\x5d\xbf\x8f\x46\xee\x13\x89\xc1\x80\xc0\xea\x8a\x0a\xec\xc3\xbb\xff\x6b\xa6\x4e\x31\x5e\xbd\x4b\x5b\x08\xb2\x87\xb0\x23\xb6\xc5\xf8\x45\x43\x5b\x7c\x69\xe1\x6c\x06\xa0\x1d\xc6\xe6\xee\x6e\xc5\xd1\xa8\x9a\xb1\x9b\xfb\xbc\x36\xc5\x0c\x97\x15\xc9\x8b\x6d\x03\xad\x7a\xbd\xd1\x56\x70\x5b\x9b\xbd\x1e\x67\x79\x9d\xa5\x14\xdb\xc4\x07\xdd\xde\x72\x1d\x92\x7a\xc7\xe8\x1a\x9a\xe7\x74\x13\x87\xb4\x1f\xd2\x02\x2b\x59\x44\xcf\x09\xcb\xb7\xbb\xf1\xfc\x96\xb6\x96\x9b\x9c\xb0\xfa\x6f\x00\x00\x00\xff\xff\xf5\x66\x3a\xde\x76\x34\x00\x00"),
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
