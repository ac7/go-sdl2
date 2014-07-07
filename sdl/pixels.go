package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixels uint8
	BytesPerPixel uint8
	padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	RefCount      int
	Next          *PixelFormat
}

type Palette struct {
	Ncolors  int
	Colors   *Color
	Version  uint32
	RefCount int
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c Color) Uint32() uint32 {
	var v uint32
	v |= uint32(c.R) << 16
	v |= uint32(c.G) << 8
	v |= uint32(c.B)
	return v
}

const (
	PIXELTYPE_UNKNOWN  = C.SDL_PIXELTYPE_UNKNOWN
	PIXELTYPE_INDEX1   = C.SDL_PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4   = C.SDL_PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8   = C.SDL_PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8  = C.SDL_PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16 = C.SDL_PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32 = C.SDL_PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8  = C.SDL_PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16 = C.SDL_PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32 = C.SDL_PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16 = C.SDL_PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32 = C.SDL_PIXELTYPE_ARRAYF32
)

/** Bitmap pixel order high bit -> low bit. */
const (
	BITMAPORDER_NONE = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 = C.SDL_BITMAPORDER_1234
)

/** Packed component order high bit -> low bit. */
const (
	PACKEDORDER_NONE = C.SDL_PACKEDORDER_NONE
	PACKEDORDER_XRGB = C.SDL_PACKEDORDER_XRGB
	PACKEDORDER_RGBX = C.SDL_PACKEDORDER_RGBX
	PACKEDORDER_ARGB = C.SDL_PACKEDORDER_ARGB
	PACKEDORDER_RGBA = C.SDL_PACKEDORDER_RGBA
	PACKEDORDER_XBGR = C.SDL_PACKEDORDER_XBGR
	PACKEDORDER_BGRX = C.SDL_PACKEDORDER_BGRX
	PACKEDORDER_ABGR = C.SDL_PACKEDORDER_ABGR
	PACKEDORDER_BGRA = C.SDL_PACKEDORDER_BGRA
)

/** Array component order low byte -> high byte. */
const (
	ARRAYORDER_NONE = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR = C.SDL_ARRAYORDER_ABGR
)

/** Packed component layout. */
const (
	PACKEDLAYOUT_NONE    = C.SDL_PACKEDLAYOUT_NONE
	PACKEDLAYOUT_332     = C.SDL_PACKEDLAYOUT_332
	PACKEDLAYOUT_4444    = C.SDL_PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555    = C.SDL_PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551    = C.SDL_PACKEDLAYOUT_5551
	PACKEDLAYOUT_565     = C.SDL_PACKEDLAYOUT_565
	PACKEDLAYOUT_8888    = C.SDL_PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010 = C.SDL_PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102 = C.SDL_PACKEDLAYOUT_1010102
)

const (
	PIXELFORMAT_UNKNOWN     = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_RGB444      = C.SDL_PIXELFORMAT_RGB444
	PIXELFORMAT_RGB555      = C.SDL_PIXELFORMAT_RGB555
	PIXELFORMAT_BGR555      = C.SDL_PIXELFORMAT_BGR555
	PIXELFORMAT_ARGB4444    = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444    = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444    = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444    = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555    = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551    = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555    = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551    = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_RGB888      = C.SDL_PIXELFORMAT_RGB888
	PIXELFORMAT_RGBX8888    = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_BGR888      = C.SDL_PIXELFORMAT_BGR888
	PIXELFORMAT_BGRX8888    = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888    = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888    = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888    = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888    = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_ARGB2101010 = C.SDL_PIXELFORMAT_ARGB2101010
	PIXELFORMAT_YV12        = C.SDL_PIXELFORMAT_YV12
	PIXELFORMAT_IYUV        = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2        = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY        = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU        = C.SDL_PIXELFORMAT_YVYU
)

func (fmt *PixelFormat) cptr() *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(fmt))
}

func (p *Palette) cptr() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(p))
}
