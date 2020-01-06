// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

/*
Objects:

Texture
Buffer
FrameBuffer
RenderBuffer
Program
UniformLocation
Shader

PerFragment:

void blendColor(float red, float green, float blue, float alpha)
void blendEquation(enum mode)
void blendEquationSeparate(enum modeRGB, enum modeAlpha)
void blendFunc(enum sfactor, enum dfactor)
void blendFuncSeparate(enum srcRGB, enum dstRGB, enum srcAlpha, enum dstAlpha)
void depthFunc(enum func)
void sampleCoverage(float value, bool invert)
void stencilFunc(int func, int ref, uint mask)
void stencilFuncSeparate(enum face, enum func, int ref, uint mask)
void stencilOp(enum fail, enum zfail, enum zpass)
void stencilOpSeparate(enum face, enum fail, enum zfail, enum zpass)

FrameBuffer:

void clear(ulong mask)
void clearColor(float red, float green, float blue, float alpha)
void clearDepth(float depth)
void clearStencil(int s)
void colorMask(bool red, bool green, bool blue, bool alpha)
void depthMask(bool flag)
void stencilMask(uint mask)
void stencilMaskSeparate(enum face, uint mask)
void bindFramebuffer(enum target, Object framebuffer)
enum checkFramebufferStatus(enum target)
Object createFramebuffer()
void deleteFramebuffer(Object buffer)
void framebufferRenderbuffer(enum target, enum attachment, enum renderbuffertarget, Object renderbuffer)
bool isFramebuffer(Object framebuffer)
void framebufferTexture2D(enum target, enum attachment, enum textarget, Object texture, int level)
any getFramebufferAttachmentParameter(enum target, enum attachment, enum pname)

Buffer:

void bindBuffer(enum target, Object buffer)
void bufferData(enum target, long size, enum usage)
void bufferData(enum target, Object data, enum usage)
void bufferSubData(enum target, long offset, Object data)
Object createBuffer()
void deleteBuffer(Object buffer)
any getBufferParameter(enum target, enum pname)
bool isBuffer(Object buffer)

View:

void depthRange(float zNear, float zFar)
void scissor(int x, int y, long width, long height)
void viewport(int x, int y, long width, long height)

Rasterization:

void cullFace(enum mode)
void frontFace(enum mode)
void lineWidth(float width)
void polygonOffset(float factor, float units)

Shaders:

void attachShader(Object program, Object shader)
void bindAttribLocation(Object program, uint index, string name)
void compileShader(Object shader)
Object createProgram()
Object createShader(enum type)
void deleteProgram(Object program)
void deleteShader(Object shader)
void detachShader(Object program, Object shader)
Object[ ] getAttachedShaders(Object program)
any getProgramParameter(Object program, enum pname)
string getProgramInfoLog(Object program)
any getShaderParameter(Object shader, enum pname)
string getShaderInfoLog(Object shader)
string getShaderSource(Object shader)
bool isProgram(Object program)
bool isShader(Object shader)
void linkProgram(Object program)
void shaderSource(Object shader, string source)
void useProgram(Object program)
void validateProgram(Object program)

Textures:

void activeTexture(enum texture)
void bindTexture(enum target, Object texture)
void copyTexImage2D(enum target, int level, enum internalformat, int x, int y, long width, long height, int border)
void copyTexSubImage2D(enum target, int level, int xoffset, int yoffset, int x, int y, long width, long height)
Object createTexture()
void deleteTexture(Object texture)
void generateMipmap(enum target)
any getTexParameter(enum target, enum pname)
bool isTexture(Object texture)
void texImage2D(enum target, int level, enum internalformat, long width, long height, int border, enum format, enum type, Object pixels)
void texImage2D(enum target, int level, enum internalformat, enum format, enum type, Object object)
void texParameterf(enum target, enum pname, float param)
void texParameteri(enum target, enum pname, int param)
void texSubImage2D(enum target, int level, int xoffset, int yoffset, long width, long height, enum format, enum type, Object pixels)
void texSubImage2D(enum target, int level, int xoffset, int yoffset, enum format, enum type, Object object)


Special:

void disable(enum cap)
void enable(enum cap)
void finish()
void flush()
enum getError()
any getParameter(enum pname)
void hint(enum target, enum mode)
bool isEnabled(enum cap)
void pixelStorei(enum pname, int param)

Uniforms and Attributes:

void disableVertexAttribArray(uint index)
void enableVertexAttribArray(uint index)
Object getActiveAttrib(Object program, uint index)
Object getActiveUniform(Object program, uint index)
ulong getAttribLocation(Object program, string name)
any getUniform(Object program, uint location)
uint getUniformLocation(Object program, string name)
any getVertexAttrib(uint index, enum pname)
long getVertexAttribOffset(uint index, enum pname)
void uniform[1234][fi](uint location, ...)
void uniform[1234][fi]v(uint location, Array value)
void uniformMatrix[234]fv(uint location, bool transpose, Array)
void vertexAttrib[1234]f(uint index, ...)
void vertexAttrib[1234]fv(uint index, Array value)
void vertexAttribPointer(uint index, int size, enum type, bool normalized, long stride, long offset)

RenderBuffer:

void bindRenderbuffer(enum target, Object renderbuffer)
Object createRenderbuffer()
void deleteRenderbuffer(Object renderbuffer)
any getRenderbufferParameter(enum target, enum pname)
bool isRenderbuffer(Object renderbuffer)
void renderbufferStorage(enum target, enum internalformat, long width, long height)

DrawBuffer:

void drawArrays(enum mode, int first, long count)
void drawElements(enum mode, long count, enum type, long offset)

ReadPixels:

void readPixels(int x, int y, long width, long height, enum format, enum type, Object pixels)

*/

import (
	"fmt"
	"image"
	"log"
	"reflect"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Texture struct{ uint32 }
type Buffer struct{ uint32 }
type FrameBuffer struct{ uint32 }
type RenderBuffer struct{ uint32 }
type Program struct{ uint32 }
type UniformLocation struct{ int32 }
type Shader struct{ uint32 }

type Context struct {
	ACCUM_ADJACENT_PAIRS_NV                                    int
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              int
	ACTIVE_ATTRIBUTES                                          int
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                int
	ACTIVE_PROGRAM                                             int
	ACTIVE_PROGRAM_EXT                                         int
	ACTIVE_RESOURCES                                           int
	ACTIVE_SUBROUTINES                                         int
	ACTIVE_SUBROUTINE_MAX_LENGTH                               int
	ACTIVE_SUBROUTINE_UNIFORMS                                 int
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        int
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       int
	ACTIVE_TEXTURE                                             int
	ACTIVE_UNIFORMS                                            int
	ACTIVE_UNIFORM_BLOCKS                                      int
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       int
	ACTIVE_UNIFORM_MAX_LENGTH                                  int
	ACTIVE_VARIABLES                                           int
	ADJACENT_PAIRS_NV                                          int
	AFFINE_2D_NV                                               int
	AFFINE_3D_NV                                               int
	ALIASED_LINE_WIDTH_RANGE                                   int
	ALL_BARRIER_BITS                                           int
	ALL_SHADER_BITS                                            int
	ALL_SHADER_BITS_EXT                                        int
	ALPHA                                                      int
	ALPHA_REF_COMMAND_NV                                       int
	ALREADY_SIGNALED                                           int
	ALWAYS                                                     int
	AND                                                        int
	AND_INVERTED                                               int
	AND_REVERSE                                                int
	ANY_SAMPLES_PASSED                                         int
	ANY_SAMPLES_PASSED_CONSERVATIVE                            int
	ARC_TO_NV                                                  int
	ARRAY_BUFFER                                               int
	ARRAY_BUFFER_BINDING                                       int
	ARRAY_SIZE                                                 int
	ARRAY_STRIDE                                               int
	ATOMIC_COUNTER_BARRIER_BIT                                 int
	ATOMIC_COUNTER_BUFFER                                      int
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               int
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        int
	ATOMIC_COUNTER_BUFFER_BINDING                              int
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            int
	ATOMIC_COUNTER_BUFFER_INDEX                                int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER int
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          int
	ATOMIC_COUNTER_BUFFER_SIZE                                 int
	ATOMIC_COUNTER_BUFFER_START                                int
	ATTACHED_SHADERS                                           int
	ATTRIBUTE_ADDRESS_COMMAND_NV                               int
	AUTO_GENERATE_MIPMAP                                       int
	BACK                                                       int
	BACK_LEFT                                                  int
	BACK_RIGHT                                                 int
	BEVEL_NV                                                   int
	BGR                                                        int
	BGRA                                                       int
	BGRA_INTEGER                                               int
	BGR_INTEGER                                                int
	BLACKHOLE_RENDER_INTEL                                     int
	BLEND                                                      int
	BLEND_ADVANCED_COHERENT_KHR                                int
	BLEND_ADVANCED_COHERENT_NV                                 int
	BLEND_COLOR                                                int
	BLEND_COLOR_COMMAND_NV                                     int
	BLEND_DST                                                  int
	BLEND_DST_ALPHA                                            int
	BLEND_DST_RGB                                              int
	BLEND_EQUATION                                             int
	BLEND_EQUATION_ALPHA                                       int
	BLEND_EQUATION_RGB                                         int
	BLEND_OVERLAP_NV                                           int
	BLEND_PREMULTIPLIED_SRC_NV                                 int
	BLEND_SRC                                                  int
	BLEND_SRC_ALPHA                                            int
	BLEND_SRC_RGB                                              int
	BLOCK_INDEX                                                int
	BLUE                                                       int
	BLUE_INTEGER                                               int
	BLUE_NV                                                    int
	BOLD_BIT_NV                                                int
	BOOL                                                       int
	BOOL_VEC2                                                  int
	BOOL_VEC3                                                  int
	BOOL_VEC4                                                  int
	BOUNDING_BOX_NV                                            int
	BOUNDING_BOX_OF_BOUNDING_BOXES_NV                          int
	BUFFER                                                     int
	BUFFER_ACCESS                                              int
	BUFFER_ACCESS_FLAGS                                        int
	BUFFER_BINDING                                             int
	BUFFER_DATA_SIZE                                           int
	BUFFER_GPU_ADDRESS_NV                                      int
	BUFFER_IMMUTABLE_STORAGE                                   int
	BUFFER_KHR                                                 int
	BUFFER_MAPPED                                              int
	BUFFER_MAP_LENGTH                                          int
	BUFFER_MAP_OFFSET                                          int
	BUFFER_MAP_POINTER                                         int
	BUFFER_OBJECT_EXT                                          int
	BUFFER_SIZE                                                int
	BUFFER_STORAGE_FLAGS                                       int
	BUFFER_UPDATE_BARRIER_BIT                                  int
	BUFFER_USAGE                                               int
	BUFFER_VARIABLE                                            int
	BYTE                                                       int
	CAVEAT_SUPPORT                                             int
	CCW                                                        int
	CIRCULAR_CCW_ARC_TO_NV                                     int
	CIRCULAR_CW_ARC_TO_NV                                      int
	CIRCULAR_TANGENT_ARC_TO_NV                                 int
	CLAMP_READ_COLOR                                           int
	CLAMP_TO_BORDER                                            int
	CLAMP_TO_BORDER_ARB                                        int
	CLAMP_TO_EDGE                                              int
	CLEAR                                                      int
	CLEAR_BUFFER                                               int
	CLEAR_TEXTURE                                              int
	CLIENT_MAPPED_BUFFER_BARRIER_BIT                           int
	CLIENT_STORAGE_BIT                                         int
	CLIPPING_INPUT_PRIMITIVES                                  int
	CLIPPING_INPUT_PRIMITIVES_ARB                              int
	CLIPPING_OUTPUT_PRIMITIVES                                 int
	CLIPPING_OUTPUT_PRIMITIVES_ARB                             int
	CLIP_DEPTH_MODE                                            int
	CLIP_DISTANCE0                                             int
	CLIP_DISTANCE1                                             int
	CLIP_DISTANCE2                                             int
	CLIP_DISTANCE3                                             int
	CLIP_DISTANCE4                                             int
	CLIP_DISTANCE5                                             int
	CLIP_DISTANCE6                                             int
	CLIP_DISTANCE7                                             int
	CLIP_ORIGIN                                                int
	CLOSE_PATH_NV                                              int
	COLOR                                                      int
	COLORBURN_KHR                                              int
	COLORBURN_NV                                               int
	COLORDODGE_KHR                                             int
	COLORDODGE_NV                                              int
	COLOR_ARRAY_ADDRESS_NV                                     int
	COLOR_ARRAY_LENGTH_NV                                      int
	COLOR_ATTACHMENT0                                          int
	COLOR_ATTACHMENT1                                          int
	COLOR_ATTACHMENT10                                         int
	COLOR_ATTACHMENT11                                         int
	COLOR_ATTACHMENT12                                         int
	COLOR_ATTACHMENT13                                         int
	COLOR_ATTACHMENT14                                         int
	COLOR_ATTACHMENT15                                         int
	COLOR_ATTACHMENT16                                         int
	COLOR_ATTACHMENT17                                         int
	COLOR_ATTACHMENT18                                         int
	COLOR_ATTACHMENT19                                         int
	COLOR_ATTACHMENT2                                          int
	COLOR_ATTACHMENT20                                         int
	COLOR_ATTACHMENT21                                         int
	COLOR_ATTACHMENT22                                         int
	COLOR_ATTACHMENT23                                         int
	COLOR_ATTACHMENT24                                         int
	COLOR_ATTACHMENT25                                         int
	COLOR_ATTACHMENT26                                         int
	COLOR_ATTACHMENT27                                         int
	COLOR_ATTACHMENT28                                         int
	COLOR_ATTACHMENT29                                         int
	COLOR_ATTACHMENT3                                          int
	COLOR_ATTACHMENT30                                         int
	COLOR_ATTACHMENT31                                         int
	COLOR_ATTACHMENT4                                          int
	COLOR_ATTACHMENT5                                          int
	COLOR_ATTACHMENT6                                          int
	COLOR_ATTACHMENT7                                          int
	COLOR_ATTACHMENT8                                          int
	COLOR_ATTACHMENT9                                          int
	COLOR_BUFFER_BIT                                           int
	COLOR_CLEAR_VALUE                                          int
	COLOR_COMPONENTS                                           int
	COLOR_ENCODING                                             int
	COLOR_LOGIC_OP                                             int
	COLOR_RENDERABLE                                           int
	COLOR_SAMPLES_NV                                           int
	COLOR_WRITEMASK                                            int
	COMMAND_BARRIER_BIT                                        int
	COMPARE_REF_TO_TEXTURE                                     int
	COMPATIBLE_SUBROUTINES                                     int
	COMPILE_STATUS                                             uint32
	COMPLETION_STATUS_ARB                                      int
	COMPLETION_STATUS_KHR                                      int
	COMPRESSED_R11_EAC                                         int
	COMPRESSED_RED                                             int
	COMPRESSED_RED_RGTC1                                       int
	COMPRESSED_RG                                              int
	COMPRESSED_RG11_EAC                                        int
	COMPRESSED_RGB                                             int
	COMPRESSED_RGB8_ETC2                                       int
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   int
	COMPRESSED_RGBA                                            int
	COMPRESSED_RGBA8_ETC2_EAC                                  int
	COMPRESSED_RGBA_ASTC_10x10_KHR                             int
	COMPRESSED_RGBA_ASTC_10x5_KHR                              int
	COMPRESSED_RGBA_ASTC_10x6_KHR                              int
	COMPRESSED_RGBA_ASTC_10x8_KHR                              int
	COMPRESSED_RGBA_ASTC_12x10_KHR                             int
	COMPRESSED_RGBA_ASTC_12x12_KHR                             int
	COMPRESSED_RGBA_ASTC_4x4_KHR                               int
	COMPRESSED_RGBA_ASTC_5x4_KHR                               int
	COMPRESSED_RGBA_ASTC_5x5_KHR                               int
	COMPRESSED_RGBA_ASTC_6x5_KHR                               int
	COMPRESSED_RGBA_ASTC_6x6_KHR                               int
	COMPRESSED_RGBA_ASTC_8x5_KHR                               int
	COMPRESSED_RGBA_ASTC_8x6_KHR                               int
	COMPRESSED_RGBA_ASTC_8x8_KHR                               int
	COMPRESSED_RGBA_BPTC_UNORM                                 int
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             int
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              int
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              int
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              int
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT                           int
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       int
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT                         int
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     int
	COMPRESSED_RGB_S3TC_DXT1_EXT                               int
	COMPRESSED_RG_RGTC2                                        int
	COMPRESSED_SIGNED_R11_EAC                                  int
	COMPRESSED_SIGNED_RED_RGTC1                                int
	COMPRESSED_SIGNED_RG11_EAC                                 int
	COMPRESSED_SIGNED_RG_RGTC2                                 int
	COMPRESSED_SRGB                                            int
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     int
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      int
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      int
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      int
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     int
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     int
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       int
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           int
	COMPRESSED_SRGB8_ETC2                                      int
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  int
	COMPRESSED_SRGB_ALPHA                                      int
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM                           int
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       int
	COMPRESSED_TEXTURE_FORMATS                                 int
	COMPUTE_SHADER                                             int
	COMPUTE_SHADER_BIT                                         int
	COMPUTE_SHADER_INVOCATIONS                                 int
	COMPUTE_SHADER_INVOCATIONS_ARB                             int
	COMPUTE_SUBROUTINE                                         int
	COMPUTE_SUBROUTINE_UNIFORM                                 int
	COMPUTE_TEXTURE                                            int
	COMPUTE_WORK_GROUP_SIZE                                    int
	CONDITION_SATISFIED                                        int
	CONFORMANT_NV                                              int
	CONIC_CURVE_TO_NV                                          int
	CONJOINT_NV                                                int
	CONSERVATIVE_RASTERIZATION_INTEL                           int
	CONSERVATIVE_RASTERIZATION_NV                              int
	CONSERVATIVE_RASTER_DILATE_GRANULARITY_NV                  int
	CONSERVATIVE_RASTER_DILATE_NV                              int
	CONSERVATIVE_RASTER_DILATE_RANGE_NV                        int
	CONSERVATIVE_RASTER_MODE_NV                                int
	CONSERVATIVE_RASTER_MODE_POST_SNAP_NV                      int
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_NV                       int
	CONSERVATIVE_RASTER_MODE_PRE_SNAP_TRIANGLES_NV             int
	CONSTANT_ALPHA                                             int
	CONSTANT_COLOR                                             int
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          int
	CONTEXT_CORE_PROFILE_BIT                                   int
	CONTEXT_FLAGS                                              int
	CONTEXT_FLAG_DEBUG_BIT                                     int
	CONTEXT_FLAG_DEBUG_BIT_KHR                                 int
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        int
	CONTEXT_FLAG_NO_ERROR_BIT                                  int
	CONTEXT_FLAG_NO_ERROR_BIT_KHR                              int
	CONTEXT_FLAG_ROBUST_ACCESS_BIT                             int
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         int
	CONTEXT_LOST                                               int
	CONTEXT_LOST_KHR                                           int
	CONTEXT_PROFILE_MASK                                       int
	CONTEXT_RELEASE_BEHAVIOR                                   int
	CONTEXT_RELEASE_BEHAVIOR_FLUSH                             int
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR                         int
	CONTEXT_RELEASE_BEHAVIOR_KHR                               int
	CONTEXT_ROBUST_ACCESS                                      int
	CONTEXT_ROBUST_ACCESS_KHR                                  int
	CONTRAST_NV                                                int
	CONVEX_HULL_NV                                             int
	COPY                                                       int
	COPY_INVERTED                                              int
	COPY_READ_BUFFER                                           int
	COPY_READ_BUFFER_BINDING                                   int
	COPY_WRITE_BUFFER                                          int
	COPY_WRITE_BUFFER_BINDING                                  int
	COUNTER_RANGE_AMD                                          int
	COUNTER_TYPE_AMD                                           int
	COUNT_DOWN_NV                                              int
	COUNT_UP_NV                                                int
	COVERAGE_MODULATION_NV                                     int
	COVERAGE_MODULATION_TABLE_NV                               int
	COVERAGE_MODULATION_TABLE_SIZE_NV                          int
	CUBIC_CURVE_TO_NV                                          int
	CULL_FACE                                                  int
	CULL_FACE_MODE                                             int
	CURRENT_PROGRAM                                            int
	CURRENT_QUERY                                              int
	CURRENT_VERTEX_ATTRIB                                      int
	CW                                                         int
	DARKEN_KHR                                                 int
	DARKEN_NV                                                  int
	DEBUG_CALLBACK_FUNCTION                                    int
	DEBUG_CALLBACK_FUNCTION_ARB                                int
	DEBUG_CALLBACK_FUNCTION_KHR                                int
	DEBUG_CALLBACK_USER_PARAM                                  int
	DEBUG_CALLBACK_USER_PARAM_ARB                              int
	DEBUG_CALLBACK_USER_PARAM_KHR                              int
	DEBUG_GROUP_STACK_DEPTH                                    int
	DEBUG_GROUP_STACK_DEPTH_KHR                                int
	DEBUG_LOGGED_MESSAGES                                      int
	DEBUG_LOGGED_MESSAGES_ARB                                  int
	DEBUG_LOGGED_MESSAGES_KHR                                  int
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           int
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       int
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR                       int
	DEBUG_OUTPUT                                               int
	DEBUG_OUTPUT_KHR                                           int
	DEBUG_OUTPUT_SYNCHRONOUS                                   int
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               int
	DEBUG_OUTPUT_SYNCHRONOUS_KHR                               int
	DEBUG_SEVERITY_HIGH                                        int
	DEBUG_SEVERITY_HIGH_ARB                                    int
	DEBUG_SEVERITY_HIGH_KHR                                    int
	DEBUG_SEVERITY_LOW                                         int
	DEBUG_SEVERITY_LOW_ARB                                     int
	DEBUG_SEVERITY_LOW_KHR                                     int
	DEBUG_SEVERITY_MEDIUM                                      int
	DEBUG_SEVERITY_MEDIUM_ARB                                  int
	DEBUG_SEVERITY_MEDIUM_KHR                                  int
	DEBUG_SEVERITY_NOTIFICATION                                int
	DEBUG_SEVERITY_NOTIFICATION_KHR                            int
	DEBUG_SOURCE_API                                           int
	DEBUG_SOURCE_API_ARB                                       int
	DEBUG_SOURCE_API_KHR                                       int
	DEBUG_SOURCE_APPLICATION                                   int
	DEBUG_SOURCE_APPLICATION_ARB                               int
	DEBUG_SOURCE_APPLICATION_KHR                               int
	DEBUG_SOURCE_OTHER                                         int
	DEBUG_SOURCE_OTHER_ARB                                     int
	DEBUG_SOURCE_OTHER_KHR                                     int
	DEBUG_SOURCE_SHADER_COMPILER                               int
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           int
	DEBUG_SOURCE_SHADER_COMPILER_KHR                           int
	DEBUG_SOURCE_THIRD_PARTY                                   int
	DEBUG_SOURCE_THIRD_PARTY_ARB                               int
	DEBUG_SOURCE_THIRD_PARTY_KHR                               int
	DEBUG_SOURCE_WINDOW_SYSTEM                                 int
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             int
	DEBUG_SOURCE_WINDOW_SYSTEM_KHR                             int
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             int
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         int
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR                         int
	DEBUG_TYPE_ERROR                                           int
	DEBUG_TYPE_ERROR_ARB                                       int
	DEBUG_TYPE_ERROR_KHR                                       int
	DEBUG_TYPE_MARKER                                          int
	DEBUG_TYPE_MARKER_KHR                                      int
	DEBUG_TYPE_OTHER                                           int
	DEBUG_TYPE_OTHER_ARB                                       int
	DEBUG_TYPE_OTHER_KHR                                       int
	DEBUG_TYPE_PERFORMANCE                                     int
	DEBUG_TYPE_PERFORMANCE_ARB                                 int
	DEBUG_TYPE_PERFORMANCE_KHR                                 int
	DEBUG_TYPE_POP_GROUP                                       int
	DEBUG_TYPE_POP_GROUP_KHR                                   int
	DEBUG_TYPE_PORTABILITY                                     int
	DEBUG_TYPE_PORTABILITY_ARB                                 int
	DEBUG_TYPE_PORTABILITY_KHR                                 int
	DEBUG_TYPE_PUSH_GROUP                                      int
	DEBUG_TYPE_PUSH_GROUP_KHR                                  int
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              int
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          int
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR                          int
	DECODE_EXT                                                 int
	DECR                                                       int
	DECR_WRAP                                                  int
	DELETE_STATUS                                              int
	DEPTH                                                      int
	DEPTH24_STENCIL8                                           int
	DEPTH32F_STENCIL8                                          int
	DEPTH_ATTACHMENT                                           int
	DEPTH_BUFFER_BIT                                           int
	DEPTH_CLAMP                                                int
	DEPTH_CLEAR_VALUE                                          int
	DEPTH_COMPONENT                                            int
	DEPTH_COMPONENT16                                          int
	DEPTH_COMPONENT24                                          int
	DEPTH_COMPONENT32                                          int
	DEPTH_COMPONENT32F                                         int
	DEPTH_COMPONENTS                                           int
	DEPTH_FUNC                                                 int
	DEPTH_RANGE                                                int
	DEPTH_RENDERABLE                                           int
	DEPTH_SAMPLES_NV                                           int
	DEPTH_STENCIL                                              int
	DEPTH_STENCIL_ATTACHMENT                                   int
	DEPTH_STENCIL_TEXTURE_MODE                                 int
	DEPTH_TEST                                                 int
	DEPTH_WRITEMASK                                            int
	DIFFERENCE_KHR                                             int
	DIFFERENCE_NV                                              int
	DISJOINT_NV                                                int
	DISPATCH_INDIRECT_BUFFER                                   int
	DISPATCH_INDIRECT_BUFFER_BINDING                           int
	DITHER                                                     int
	DONT_CARE                                                  int
	DOUBLE                                                     int
	DOUBLEBUFFER                                               int
	DOUBLE_MAT2                                                int
	DOUBLE_MAT2x3                                              int
	DOUBLE_MAT2x4                                              int
	DOUBLE_MAT3                                                int
	DOUBLE_MAT3x2                                              int
	DOUBLE_MAT3x4                                              int
	DOUBLE_MAT4                                                int
	DOUBLE_MAT4x2                                              int
	DOUBLE_MAT4x3                                              int
	DOUBLE_VEC2                                                int
	DOUBLE_VEC3                                                int
	DOUBLE_VEC4                                                int
	DRAW_ARRAYS_COMMAND_NV                                     int
	DRAW_ARRAYS_INSTANCED_COMMAND_NV                           int
	DRAW_ARRAYS_STRIP_COMMAND_NV                               int
	DRAW_BUFFER                                                int
	DRAW_BUFFER0                                               int
	DRAW_BUFFER1                                               int
	DRAW_BUFFER10                                              int
	DRAW_BUFFER11                                              int
	DRAW_BUFFER12                                              int
	DRAW_BUFFER13                                              int
	DRAW_BUFFER14                                              int
	DRAW_BUFFER15                                              int
	DRAW_BUFFER2                                               int
	DRAW_BUFFER3                                               int
	DRAW_BUFFER4                                               int
	DRAW_BUFFER5                                               int
	DRAW_BUFFER6                                               int
	DRAW_BUFFER7                                               int
	DRAW_BUFFER8                                               int
	DRAW_BUFFER9                                               int
	DRAW_ELEMENTS_COMMAND_NV                                   int
	DRAW_ELEMENTS_INSTANCED_COMMAND_NV                         int
	DRAW_ELEMENTS_STRIP_COMMAND_NV                             int
	DRAW_FRAMEBUFFER                                           int
	DRAW_FRAMEBUFFER_BINDING                                   int
	DRAW_INDIRECT_ADDRESS_NV                                   int
	DRAW_INDIRECT_BUFFER                                       int
	DRAW_INDIRECT_BUFFER_BINDING                               int
	DRAW_INDIRECT_LENGTH_NV                                    int
	DRAW_INDIRECT_UNIFIED_NV                                   int
	DST_ALPHA                                                  int
	DST_ATOP_NV                                                int
	DST_COLOR                                                  int
	DST_IN_NV                                                  int
	DST_NV                                                     int
	DST_OUT_NV                                                 int
	DST_OVER_NV                                                int
	DUP_FIRST_CUBIC_CURVE_TO_NV                                int
	DUP_LAST_CUBIC_CURVE_TO_NV                                 int
	DYNAMIC_COPY                                               int
	DYNAMIC_DRAW                                               int
	DYNAMIC_READ                                               int
	DYNAMIC_STORAGE_BIT                                        int
	EDGE_FLAG_ARRAY_ADDRESS_NV                                 int
	EDGE_FLAG_ARRAY_LENGTH_NV                                  int
	EFFECTIVE_RASTER_SAMPLES_EXT                               int
	ELEMENT_ADDRESS_COMMAND_NV                                 int
	ELEMENT_ARRAY_ADDRESS_NV                                   int
	ELEMENT_ARRAY_BARRIER_BIT                                  int
	ELEMENT_ARRAY_BUFFER                                       int
	ELEMENT_ARRAY_BUFFER_BINDING                               int
	ELEMENT_ARRAY_LENGTH_NV                                    int
	ELEMENT_ARRAY_UNIFIED_NV                                   int
	EQUAL                                                      int
	EQUIV                                                      int
	EXCLUSION_KHR                                              int
	EXCLUSION_NV                                               int
	EXCLUSIVE_EXT                                              int
	EXTENSIONS                                                 int
	FACTOR_MAX_AMD                                             int
	FACTOR_MIN_AMD                                             int
	FALSE                                                      int
	FASTEST                                                    int
	FILE_NAME_NV                                               int
	FILL                                                       int
	FILL_RECTANGLE_NV                                          int
	FILTER                                                     int
	FIRST_TO_REST_NV                                           int
	FIRST_VERTEX_CONVENTION                                    int
	FIXED                                                      int
	FIXED_ONLY                                                 int
	FLOAT                                                      int
	FLOAT16_NV                                                 int
	FLOAT16_VEC2_NV                                            int
	FLOAT16_VEC3_NV                                            int
	FLOAT16_VEC4_NV                                            int
	FLOAT_32_UNSIGNED_INT_24_8_REV                             int
	FLOAT_MAT2                                                 int
	FLOAT_MAT2x3                                               int
	FLOAT_MAT2x4                                               int
	FLOAT_MAT3                                                 int
	FLOAT_MAT3x2                                               int
	FLOAT_MAT3x4                                               int
	FLOAT_MAT4                                                 int
	FLOAT_MAT4x2                                               int
	FLOAT_MAT4x3                                               int
	FLOAT_VEC2                                                 int
	FLOAT_VEC3                                                 int
	FLOAT_VEC4                                                 int
	FOG_COORD_ARRAY_ADDRESS_NV                                 int
	FOG_COORD_ARRAY_LENGTH_NV                                  int
	FONT_ASCENDER_BIT_NV                                       int
	FONT_DESCENDER_BIT_NV                                      int
	FONT_GLYPHS_AVAILABLE_NV                                   int
	FONT_HAS_KERNING_BIT_NV                                    int
	FONT_HEIGHT_BIT_NV                                         int
	FONT_MAX_ADVANCE_HEIGHT_BIT_NV                             int
	FONT_MAX_ADVANCE_WIDTH_BIT_NV                              int
	FONT_NUM_GLYPH_INDICES_BIT_NV                              int
	FONT_TARGET_UNAVAILABLE_NV                                 int
	FONT_UNAVAILABLE_NV                                        int
	FONT_UNDERLINE_POSITION_BIT_NV                             int
	FONT_UNDERLINE_THICKNESS_BIT_NV                            int
	FONT_UNINTELLIGIBLE_NV                                     int
	FONT_UNITS_PER_EM_BIT_NV                                   int
	FONT_X_MAX_BOUNDS_BIT_NV                                   int
	FONT_X_MIN_BOUNDS_BIT_NV                                   int
	FONT_Y_MAX_BOUNDS_BIT_NV                                   int
	FONT_Y_MIN_BOUNDS_BIT_NV                                   int
	FRACTIONAL_EVEN                                            int
	FRACTIONAL_ODD                                             int
	FRAGMENT_COVERAGE_COLOR_NV                                 int
	FRAGMENT_COVERAGE_TO_COLOR_NV                              int
	FRAGMENT_INPUT_NV                                          int
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         int
	FRAGMENT_SHADER                                            int
	FRAGMENT_SHADER_BIT                                        int
	FRAGMENT_SHADER_BIT_EXT                                    int
	FRAGMENT_SHADER_DERIVATIVE_HINT                            int
	FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT                       int
	FRAGMENT_SHADER_INVOCATIONS                                int
	FRAGMENT_SHADER_INVOCATIONS_ARB                            int
	FRAGMENT_SUBROUTINE                                        int
	FRAGMENT_SUBROUTINE_UNIFORM                                int
	FRAGMENT_TEXTURE                                           int
	FRAMEBUFFER                                                int
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          int
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           int
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      int
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      int
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          int
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          int
	FRAMEBUFFER_ATTACHMENT_LAYERED                             int
	FRAMEBUFFER_ATTACHMENT_LAYERED_ARB                         int
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         int
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         int
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            int
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR         int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       int
	FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR               int
	FRAMEBUFFER_BARRIER_BIT                                    int
	FRAMEBUFFER_BINDING                                        int
	FRAMEBUFFER_BLEND                                          int
	FRAMEBUFFER_COMPLETE                                       int
	FRAMEBUFFER_DEFAULT                                        int
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 int
	FRAMEBUFFER_DEFAULT_HEIGHT                                 int
	FRAMEBUFFER_DEFAULT_LAYERS                                 int
	FRAMEBUFFER_DEFAULT_SAMPLES                                int
	FRAMEBUFFER_DEFAULT_WIDTH                                  int
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          int
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         int
	FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB                     int
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       int
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB                   int
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  int
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         int
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         int
	FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR                    int
	FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_ARB              int
	FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_NV               int
	FRAMEBUFFER_RENDERABLE                                     int
	FRAMEBUFFER_RENDERABLE_LAYERED                             int
	FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_ARB                 int
	FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_NV                  int
	FRAMEBUFFER_SRGB                                           int
	FRAMEBUFFER_UNDEFINED                                      int
	FRAMEBUFFER_UNSUPPORTED                                    int
	FRONT                                                      int
	FRONT_AND_BACK                                             int
	FRONT_FACE                                                 int
	FRONT_FACE_COMMAND_NV                                      int
	FRONT_LEFT                                                 int
	FRONT_RIGHT                                                int
	FULL_SUPPORT                                               int
	FUNC_ADD                                                   int
	FUNC_REVERSE_SUBTRACT                                      int
	FUNC_SUBTRACT                                              int
	GEOMETRY_INPUT_TYPE                                        int
	GEOMETRY_INPUT_TYPE_ARB                                    int
	GEOMETRY_OUTPUT_TYPE                                       int
	GEOMETRY_OUTPUT_TYPE_ARB                                   int
	GEOMETRY_SHADER                                            int
	GEOMETRY_SHADER_ARB                                        int
	GEOMETRY_SHADER_BIT                                        int
	GEOMETRY_SHADER_INVOCATIONS                                int
	GEOMETRY_SHADER_PRIMITIVES_EMITTED                         int
	GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB                     int
	GEOMETRY_SUBROUTINE                                        int
	GEOMETRY_SUBROUTINE_UNIFORM                                int
	GEOMETRY_TEXTURE                                           int
	GEOMETRY_VERTICES_OUT                                      int
	GEOMETRY_VERTICES_OUT_ARB                                  int
	GEQUAL                                                     int
	GET_TEXTURE_IMAGE_FORMAT                                   int
	GET_TEXTURE_IMAGE_TYPE                                     int
	GLYPH_HAS_KERNING_BIT_NV                                   int
	GLYPH_HEIGHT_BIT_NV                                        int
	GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV                    int
	GLYPH_HORIZONTAL_BEARING_X_BIT_NV                          int
	GLYPH_HORIZONTAL_BEARING_Y_BIT_NV                          int
	GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV                      int
	GLYPH_VERTICAL_BEARING_X_BIT_NV                            int
	GLYPH_VERTICAL_BEARING_Y_BIT_NV                            int
	GLYPH_WIDTH_BIT_NV                                         int
	GPU_ADDRESS_NV                                             int
	GREATER                                                    int
	GREEN                                                      int
	GREEN_INTEGER                                              int
	GREEN_NV                                                   int
	GUILTY_CONTEXT_RESET                                       int
	GUILTY_CONTEXT_RESET_ARB                                   int
	GUILTY_CONTEXT_RESET_KHR                                   int
	HALF_FLOAT                                                 int
	HARDLIGHT_KHR                                              int
	HARDLIGHT_NV                                               int
	HARDMIX_NV                                                 int
	HIGH_FLOAT                                                 int
	HIGH_INT                                                   int
	HORIZONTAL_LINE_TO_NV                                      int
	HSL_COLOR_KHR                                              int
	HSL_COLOR_NV                                               int
	HSL_HUE_KHR                                                int
	HSL_HUE_NV                                                 int
	HSL_LUMINOSITY_KHR                                         int
	HSL_LUMINOSITY_NV                                          int
	HSL_SATURATION_KHR                                         int
	HSL_SATURATION_NV                                          int
	IMAGE_1D                                                   int
	IMAGE_1D_ARRAY                                             int
	IMAGE_2D                                                   int
	IMAGE_2D_ARRAY                                             int
	IMAGE_2D_MULTISAMPLE                                       int
	IMAGE_2D_MULTISAMPLE_ARRAY                                 int
	IMAGE_2D_RECT                                              int
	IMAGE_3D                                                   int
	IMAGE_BINDING_ACCESS                                       int
	IMAGE_BINDING_FORMAT                                       int
	IMAGE_BINDING_LAYER                                        int
	IMAGE_BINDING_LAYERED                                      int
	IMAGE_BINDING_LEVEL                                        int
	IMAGE_BINDING_NAME                                         int
	IMAGE_BUFFER                                               int
	IMAGE_CLASS_10_10_10_2                                     int
	IMAGE_CLASS_11_11_10                                       int
	IMAGE_CLASS_1_X_16                                         int
	IMAGE_CLASS_1_X_32                                         int
	IMAGE_CLASS_1_X_8                                          int
	IMAGE_CLASS_2_X_16                                         int
	IMAGE_CLASS_2_X_32                                         int
	IMAGE_CLASS_2_X_8                                          int
	IMAGE_CLASS_4_X_16                                         int
	IMAGE_CLASS_4_X_32                                         int
	IMAGE_CLASS_4_X_8                                          int
	IMAGE_COMPATIBILITY_CLASS                                  int
	IMAGE_CUBE                                                 int
	IMAGE_CUBE_MAP_ARRAY                                       int
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        int
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         int
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            int
	IMAGE_PIXEL_FORMAT                                         int
	IMAGE_PIXEL_TYPE                                           int
	IMAGE_TEXEL_SIZE                                           int
	IMPLEMENTATION_COLOR_READ_FORMAT                           int
	IMPLEMENTATION_COLOR_READ_TYPE                             int
	INCLUSIVE_EXT                                              int
	INCR                                                       int
	INCR_WRAP                                                  int
	INDEX_ARRAY_ADDRESS_NV                                     int
	INDEX_ARRAY_LENGTH_NV                                      int
	INFO_LOG_LENGTH                                            int
	INNOCENT_CONTEXT_RESET                                     int
	INNOCENT_CONTEXT_RESET_ARB                                 int
	INNOCENT_CONTEXT_RESET_KHR                                 int
	INT                                                        int
	INT16_NV                                                   int
	INT16_VEC2_NV                                              int
	INT16_VEC3_NV                                              int
	INT16_VEC4_NV                                              int
	INT64_ARB                                                  int
	INT64_NV                                                   int
	INT64_VEC2_ARB                                             int
	INT64_VEC2_NV                                              int
	INT64_VEC3_ARB                                             int
	INT64_VEC3_NV                                              int
	INT64_VEC4_ARB                                             int
	INT64_VEC4_NV                                              int
	INT8_NV                                                    int
	INT8_VEC2_NV                                               int
	INT8_VEC3_NV                                               int
	INT8_VEC4_NV                                               int
	INTERLEAVED_ATTRIBS                                        int
	INTERNALFORMAT_ALPHA_SIZE                                  int
	INTERNALFORMAT_ALPHA_TYPE                                  int
	INTERNALFORMAT_BLUE_SIZE                                   int
	INTERNALFORMAT_BLUE_TYPE                                   int
	INTERNALFORMAT_DEPTH_SIZE                                  int
	INTERNALFORMAT_DEPTH_TYPE                                  int
	INTERNALFORMAT_GREEN_SIZE                                  int
	INTERNALFORMAT_GREEN_TYPE                                  int
	INTERNALFORMAT_PREFERRED                                   int
	INTERNALFORMAT_RED_SIZE                                    int
	INTERNALFORMAT_RED_TYPE                                    int
	INTERNALFORMAT_SHARED_SIZE                                 int
	INTERNALFORMAT_STENCIL_SIZE                                int
	INTERNALFORMAT_STENCIL_TYPE                                int
	INTERNALFORMAT_SUPPORTED                                   int
	INT_2_10_10_10_REV                                         int
	INT_IMAGE_1D                                               int
	INT_IMAGE_1D_ARRAY                                         int
	INT_IMAGE_2D                                               int
	INT_IMAGE_2D_ARRAY                                         int
	INT_IMAGE_2D_MULTISAMPLE                                   int
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             int
	INT_IMAGE_2D_RECT                                          int
	INT_IMAGE_3D                                               int
	INT_IMAGE_BUFFER                                           int
	INT_IMAGE_CUBE                                             int
	INT_IMAGE_CUBE_MAP_ARRAY                                   int
	INT_SAMPLER_1D                                             int
	INT_SAMPLER_1D_ARRAY                                       int
	INT_SAMPLER_2D                                             int
	INT_SAMPLER_2D_ARRAY                                       int
	INT_SAMPLER_2D_MULTISAMPLE                                 int
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           int
	INT_SAMPLER_2D_RECT                                        int
	INT_SAMPLER_3D                                             int
	INT_SAMPLER_BUFFER                                         int
	INT_SAMPLER_CUBE                                           int
	INT_SAMPLER_CUBE_MAP_ARRAY                                 int
	INT_SAMPLER_CUBE_MAP_ARRAY_ARB                             int
	INT_VEC2                                                   int
	INT_VEC3                                                   int
	INT_VEC4                                                   int
	INVALID_ENUM                                               int
	INVALID_FRAMEBUFFER_OPERATION                              int
	INVALID_INDEX                                              int
	INVALID_OPERATION                                          int
	INVALID_VALUE                                              int
	INVERT                                                     int
	INVERT_OVG_NV                                              int
	INVERT_RGB_NV                                              int
	ISOLINES                                                   int
	IS_PER_PATCH                                               int
	IS_ROW_MAJOR                                               int
	ITALIC_BIT_NV                                              int
	KEEP                                                       int
	LARGE_CCW_ARC_TO_NV                                        int
	LARGE_CW_ARC_TO_NV                                         int
	LAST_VERTEX_CONVENTION                                     int
	LAYER_PROVOKING_VERTEX                                     int
	LEFT                                                       int
	LEQUAL                                                     int
	LESS                                                       int
	LIGHTEN_KHR                                                int
	LIGHTEN_NV                                                 int
	LINE                                                       int
	LINEAR                                                     int
	LINEARBURN_NV                                              int
	LINEARDODGE_NV                                             int
	LINEARLIGHT_NV                                             int
	LINEAR_MIPMAP_LINEAR                                       int
	LINEAR_MIPMAP_NEAREST                                      int
	LINES                                                      int
	LINES_ADJACENCY                                            int
	LINES_ADJACENCY_ARB                                        int
	LINE_LOOP                                                  int
	LINE_SMOOTH                                                int
	LINE_SMOOTH_HINT                                           int
	LINE_STRIP                                                 int
	LINE_STRIP_ADJACENCY                                       int
	LINE_STRIP_ADJACENCY_ARB                                   int
	LINE_TO_NV                                                 int
	LINE_WIDTH                                                 int
	LINE_WIDTH_COMMAND_NV                                      int
	LINE_WIDTH_GRANULARITY                                     int
	LINE_WIDTH_RANGE                                           int
	LINK_STATUS                                                int
	LOCATION                                                   int
	LOCATION_COMPONENT                                         int
	LOCATION_INDEX                                             int
	LOGIC_OP_MODE                                              int
	LOSE_CONTEXT_ON_RESET                                      int
	LOSE_CONTEXT_ON_RESET_ARB                                  int
	LOSE_CONTEXT_ON_RESET_KHR                                  int
	LOWER_LEFT                                                 int
	LOW_FLOAT                                                  int
	LOW_INT                                                    int
	MAJOR_VERSION                                              int
	MANUAL_GENERATE_MIPMAP                                     int
	MAP_COHERENT_BIT                                           int
	MAP_FLUSH_EXPLICIT_BIT                                     int
	MAP_INVALIDATE_BUFFER_BIT                                  int
	MAP_INVALIDATE_RANGE_BIT                                   int
	MAP_PERSISTENT_BIT                                         int
	MAP_READ_BIT                                               int
	MAP_UNSYNCHRONIZED_BIT                                     int
	MAP_WRITE_BIT                                              int
	MATRIX_STRIDE                                              int
	MAX                                                        int
	MAX_3D_TEXTURE_SIZE                                        int
	MAX_ARRAY_TEXTURE_LAYERS                                   int
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         int
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             int
	MAX_CLIP_DISTANCES                                         int
	MAX_COLOR_ATTACHMENTS                                      int
	MAX_COLOR_TEXTURE_SAMPLES                                  int
	MAX_COMBINED_ATOMIC_COUNTERS                               int
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        int
	MAX_COMBINED_CLIP_AND_CULL_DISTANCES                       int
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    int
	MAX_COMBINED_DIMENSIONS                                    int
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   int
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   int
	MAX_COMBINED_IMAGE_UNIFORMS                                int
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              int
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       int
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         int
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               int
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            int
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           int
	MAX_COMBINED_UNIFORM_BLOCKS                                int
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     int
	MAX_COMPUTE_ATOMIC_COUNTERS                                int
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         int
	MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB                    int
	MAX_COMPUTE_FIXED_GROUP_SIZE_ARB                           int
	MAX_COMPUTE_IMAGE_UNIFORMS                                 int
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          int
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             int
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            int
	MAX_COMPUTE_UNIFORM_BLOCKS                                 int
	MAX_COMPUTE_UNIFORM_COMPONENTS                             int
	MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB                 int
	MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB                        int
	MAX_COMPUTE_WORK_GROUP_COUNT                               int
	MAX_COMPUTE_WORK_GROUP_INVOCATIONS                         int
	MAX_COMPUTE_WORK_GROUP_SIZE                                int
	MAX_CUBE_MAP_TEXTURE_SIZE                                  int
	MAX_CULL_DISTANCES                                         int
	MAX_DEBUG_GROUP_STACK_DEPTH                                int
	MAX_DEBUG_GROUP_STACK_DEPTH_KHR                            int
	MAX_DEBUG_LOGGED_MESSAGES                                  int
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              int
	MAX_DEBUG_LOGGED_MESSAGES_KHR                              int
	MAX_DEBUG_MESSAGE_LENGTH                                   int
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               int
	MAX_DEBUG_MESSAGE_LENGTH_KHR                               int
	MAX_DEPTH                                                  int
	MAX_DEPTH_TEXTURE_SAMPLES                                  int
	MAX_DRAW_BUFFERS                                           int
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               int
	MAX_ELEMENTS_INDICES                                       int
	MAX_ELEMENTS_VERTICES                                      int
	MAX_ELEMENT_INDEX                                          int
	MAX_FRAGMENT_ATOMIC_COUNTERS                               int
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        int
	MAX_FRAGMENT_IMAGE_UNIFORMS                                int
	MAX_FRAGMENT_INPUT_COMPONENTS                              int
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          int
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         int
	MAX_FRAGMENT_UNIFORM_BLOCKS                                int
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            int
	MAX_FRAGMENT_UNIFORM_VECTORS                               int
	MAX_FRAMEBUFFER_HEIGHT                                     int
	MAX_FRAMEBUFFER_LAYERS                                     int
	MAX_FRAMEBUFFER_SAMPLES                                    int
	MAX_FRAMEBUFFER_WIDTH                                      int
	MAX_GEOMETRY_ATOMIC_COUNTERS                               int
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        int
	MAX_GEOMETRY_IMAGE_UNIFORMS                                int
	MAX_GEOMETRY_INPUT_COMPONENTS                              int
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             int
	MAX_GEOMETRY_OUTPUT_VERTICES                               int
	MAX_GEOMETRY_OUTPUT_VERTICES_ARB                           int
	MAX_GEOMETRY_SHADER_INVOCATIONS                            int
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         int
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           int
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB                       int
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       int
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB                   int
	MAX_GEOMETRY_UNIFORM_BLOCKS                                int
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            int
	MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB                        int
	MAX_GEOMETRY_VARYING_COMPONENTS_ARB                        int
	MAX_HEIGHT                                                 int
	MAX_IMAGE_SAMPLES                                          int
	MAX_IMAGE_UNITS                                            int
	MAX_INTEGER_SAMPLES                                        int
	MAX_LABEL_LENGTH                                           int
	MAX_LABEL_LENGTH_KHR                                       int
	MAX_LAYERS                                                 int
	MAX_MULTISAMPLE_COVERAGE_MODES_NV                          int
	MAX_NAME_LENGTH                                            int
	MAX_NUM_ACTIVE_VARIABLES                                   int
	MAX_NUM_COMPATIBLE_SUBROUTINES                             int
	MAX_PATCH_VERTICES                                         int
	MAX_PROGRAM_TEXEL_OFFSET                                   int
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  int
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          int
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      int
	MAX_RASTER_SAMPLES_EXT                                     int
	MAX_RECTANGLE_TEXTURE_SIZE                                 int
	MAX_RENDERBUFFER_SIZE                                      int
	MAX_SAMPLES                                                int
	MAX_SAMPLE_MASK_WORDS                                      int
	MAX_SERVER_WAIT_TIMEOUT                                    int
	MAX_SHADER_BUFFER_ADDRESS_NV                               int
	MAX_SHADER_COMPILER_THREADS_ARB                            int
	MAX_SHADER_COMPILER_THREADS_KHR                            int
	MAX_SHADER_STORAGE_BLOCK_SIZE                              int
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         int
	MAX_SPARSE_3D_TEXTURE_SIZE_ARB                             int
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB                        int
	MAX_SPARSE_TEXTURE_SIZE_ARB                                int
	MAX_SUBPIXEL_PRECISION_BIAS_BITS_NV                        int
	MAX_SUBROUTINES                                            int
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           int
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           int
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    int
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            int
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          int
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         int
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     int
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       int
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   int
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            int
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        int
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        int
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 int
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         int
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       int
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      int
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  int
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    int
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         int
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     int
	MAX_TESS_GEN_LEVEL                                         int
	MAX_TESS_PATCH_COMPONENTS                                  int
	MAX_TEXTURE_BUFFER_SIZE                                    int
	MAX_TEXTURE_BUFFER_SIZE_ARB                                int
	MAX_TEXTURE_IMAGE_UNITS                                    int
	MAX_TEXTURE_LOD_BIAS                                       int
	MAX_TEXTURE_MAX_ANISOTROPY                                 int
	MAX_TEXTURE_SIZE                                           int
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             int
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              int
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    int
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 int
	MAX_UNIFORM_BLOCK_SIZE                                     int
	MAX_UNIFORM_BUFFER_BINDINGS                                int
	MAX_UNIFORM_LOCATIONS                                      int
	MAX_VARYING_COMPONENTS                                     int
	MAX_VARYING_FLOATS                                         int
	MAX_VARYING_VECTORS                                        int
	MAX_VERTEX_ATOMIC_COUNTERS                                 int
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          int
	MAX_VERTEX_ATTRIBS                                         int
	MAX_VERTEX_ATTRIB_BINDINGS                                 int
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          int
	MAX_VERTEX_ATTRIB_STRIDE                                   int
	MAX_VERTEX_IMAGE_UNIFORMS                                  int
	MAX_VERTEX_OUTPUT_COMPONENTS                               int
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           int
	MAX_VERTEX_STREAMS                                         int
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             int
	MAX_VERTEX_UNIFORM_BLOCKS                                  int
	MAX_VERTEX_UNIFORM_COMPONENTS                              int
	MAX_VERTEX_UNIFORM_VECTORS                                 int
	MAX_VERTEX_VARYING_COMPONENTS_ARB                          int
	MAX_VIEWPORTS                                              int
	MAX_VIEWPORT_DIMS                                          int
	MAX_VIEWS_OVR                                              int
	MAX_WIDTH                                                  int
	MAX_WINDOW_RECTANGLES_EXT                                  int
	MEDIUM_FLOAT                                               int
	MEDIUM_INT                                                 int
	MIN                                                        int
	MINOR_VERSION                                              int
	MINUS_CLAMPED_NV                                           int
	MINUS_NV                                                   int
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          int
	MIN_MAP_BUFFER_ALIGNMENT                                   int
	MIN_PROGRAM_TEXEL_OFFSET                                   int
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          int
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      int
	MIN_SAMPLE_SHADING_VALUE                                   int
	MIN_SAMPLE_SHADING_VALUE_ARB                               int
	MIPMAP                                                     int
	MIRRORED_REPEAT                                            int
	MIRRORED_REPEAT_ARB                                        int
	MIRROR_CLAMP_TO_EDGE                                       int
	MITER_REVERT_NV                                            int
	MITER_TRUNCATE_NV                                          int
	MIXED_DEPTH_SAMPLES_SUPPORTED_NV                           int
	MIXED_STENCIL_SAMPLES_SUPPORTED_NV                         int
	MOVE_TO_CONTINUES_NV                                       int
	MOVE_TO_NV                                                 int
	MOVE_TO_RESETS_NV                                          int
	MULTIPLY_KHR                                               int
	MULTIPLY_NV                                                int
	MULTISAMPLE                                                int
	MULTISAMPLES_NV                                            int
	MULTISAMPLE_COVERAGE_MODES_NV                              int
	MULTISAMPLE_LINE_WIDTH_GRANULARITY_ARB                     int
	MULTISAMPLE_LINE_WIDTH_RANGE_ARB                           int
	MULTISAMPLE_RASTERIZATION_ALLOWED_EXT                      int
	NAMED_STRING_LENGTH_ARB                                    int
	NAMED_STRING_TYPE_ARB                                      int
	NAME_LENGTH                                                int
	NAND                                                       int
	NEAREST                                                    int
	NEAREST_MIPMAP_LINEAR                                      int
	NEAREST_MIPMAP_NEAREST                                     int
	NEGATIVE_ONE_TO_ONE                                        int
	NEVER                                                      int
	NICEST                                                     int
	NONE                                                       int
	NOOP                                                       int
	NOP_COMMAND_NV                                             int
	NOR                                                        int
	NORMAL_ARRAY_ADDRESS_NV                                    int
	NORMAL_ARRAY_LENGTH_NV                                     int
	NOTEQUAL                                                   int
	NO_ERROR                                                   int
	NO_RESET_NOTIFICATION                                      int
	NO_RESET_NOTIFICATION_ARB                                  int
	NO_RESET_NOTIFICATION_KHR                                  int
	NUM_ACTIVE_VARIABLES                                       int
	NUM_COMPATIBLE_SUBROUTINES                                 int
	NUM_COMPRESSED_TEXTURE_FORMATS                             int
	NUM_EXTENSIONS                                             int
	NUM_PROGRAM_BINARY_FORMATS                                 int
	NUM_SAMPLE_COUNTS                                          int
	NUM_SHADER_BINARY_FORMATS                                  int
	NUM_SHADING_LANGUAGE_VERSIONS                              int
	NUM_SPARSE_LEVELS_ARB                                      int
	NUM_SPIR_V_EXTENSIONS                                      int
	NUM_VIRTUAL_PAGE_SIZES_ARB                                 int
	NUM_WINDOW_RECTANGLES_EXT                                  int
	OBJECT_TYPE                                                int
	OFFSET                                                     int
	ONE                                                        int
	ONE_MINUS_CONSTANT_ALPHA                                   int
	ONE_MINUS_CONSTANT_COLOR                                   int
	ONE_MINUS_DST_ALPHA                                        int
	ONE_MINUS_DST_COLOR                                        int
	ONE_MINUS_SRC1_ALPHA                                       int
	ONE_MINUS_SRC1_COLOR                                       int
	ONE_MINUS_SRC_ALPHA                                        int
	ONE_MINUS_SRC_COLOR                                        int
	OR                                                         int
	OR_INVERTED                                                int
	OR_REVERSE                                                 int
	OUT_OF_MEMORY                                              int
	OVERLAY_KHR                                                int
	OVERLAY_NV                                                 int
	PACK_ALIGNMENT                                             int
	PACK_COMPRESSED_BLOCK_DEPTH                                int
	PACK_COMPRESSED_BLOCK_HEIGHT                               int
	PACK_COMPRESSED_BLOCK_SIZE                                 int
	PACK_COMPRESSED_BLOCK_WIDTH                                int
	PACK_IMAGE_HEIGHT                                          int
	PACK_LSB_FIRST                                             int
	PACK_ROW_LENGTH                                            int
	PACK_SKIP_IMAGES                                           int
	PACK_SKIP_PIXELS                                           int
	PACK_SKIP_ROWS                                             int
	PACK_SWAP_BYTES                                            int
	PARAMETER_BUFFER                                           int
	PARAMETER_BUFFER_ARB                                       int
	PARAMETER_BUFFER_BINDING                                   int
	PARAMETER_BUFFER_BINDING_ARB                               int
	PATCHES                                                    int
	PATCH_DEFAULT_INNER_LEVEL                                  int
	PATCH_DEFAULT_OUTER_LEVEL                                  int
	PATCH_VERTICES                                             int
	PATH_CLIENT_LENGTH_NV                                      int
	PATH_COMMAND_COUNT_NV                                      int
	PATH_COMPUTED_LENGTH_NV                                    int
	PATH_COORD_COUNT_NV                                        int
	PATH_COVER_DEPTH_FUNC_NV                                   int
	PATH_DASH_ARRAY_COUNT_NV                                   int
	PATH_DASH_CAPS_NV                                          int
	PATH_DASH_OFFSET_NV                                        int
	PATH_DASH_OFFSET_RESET_NV                                  int
	PATH_END_CAPS_NV                                           int
	PATH_ERROR_POSITION_NV                                     int
	PATH_FILL_BOUNDING_BOX_NV                                  int
	PATH_FILL_COVER_MODE_NV                                    int
	PATH_FILL_MASK_NV                                          int
	PATH_FILL_MODE_NV                                          int
	PATH_FORMAT_PS_NV                                          int
	PATH_FORMAT_SVG_NV                                         int
	PATH_GEN_COEFF_NV                                          int
	PATH_GEN_COMPONENTS_NV                                     int
	PATH_GEN_MODE_NV                                           int
	PATH_INITIAL_DASH_CAP_NV                                   int
	PATH_INITIAL_END_CAP_NV                                    int
	PATH_JOIN_STYLE_NV                                         int
	PATH_MAX_MODELVIEW_STACK_DEPTH_NV                          int
	PATH_MAX_PROJECTION_STACK_DEPTH_NV                         int
	PATH_MITER_LIMIT_NV                                        int
	PATH_MODELVIEW_MATRIX_NV                                   int
	PATH_MODELVIEW_NV                                          int
	PATH_MODELVIEW_STACK_DEPTH_NV                              int
	PATH_OBJECT_BOUNDING_BOX_NV                                int
	PATH_PROJECTION_MATRIX_NV                                  int
	PATH_PROJECTION_NV                                         int
	PATH_PROJECTION_STACK_DEPTH_NV                             int
	PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV                        int
	PATH_STENCIL_DEPTH_OFFSET_UNITS_NV                         int
	PATH_STENCIL_FUNC_NV                                       int
	PATH_STENCIL_REF_NV                                        int
	PATH_STENCIL_VALUE_MASK_NV                                 int
	PATH_STROKE_BOUNDING_BOX_NV                                int
	PATH_STROKE_COVER_MODE_NV                                  int
	PATH_STROKE_MASK_NV                                        int
	PATH_STROKE_WIDTH_NV                                       int
	PATH_TERMINAL_DASH_CAP_NV                                  int
	PATH_TERMINAL_END_CAP_NV                                   int
	PATH_TRANSPOSE_MODELVIEW_MATRIX_NV                         int
	PATH_TRANSPOSE_PROJECTION_MATRIX_NV                        int
	PERCENTAGE_AMD                                             int
	PERFMON_RESULT_AMD                                         int
	PERFMON_RESULT_AVAILABLE_AMD                               int
	PERFMON_RESULT_SIZE_AMD                                    int
	PERFQUERY_COUNTER_DATA_BOOL32_INTEL                        int
	PERFQUERY_COUNTER_DATA_DOUBLE_INTEL                        int
	PERFQUERY_COUNTER_DATA_FLOAT_INTEL                         int
	PERFQUERY_COUNTER_DATA_UINT32_INTEL                        int
	PERFQUERY_COUNTER_DATA_UINT64_INTEL                        int
	PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL                    int
	PERFQUERY_COUNTER_DURATION_NORM_INTEL                      int
	PERFQUERY_COUNTER_DURATION_RAW_INTEL                       int
	PERFQUERY_COUNTER_EVENT_INTEL                              int
	PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL                    int
	PERFQUERY_COUNTER_RAW_INTEL                                int
	PERFQUERY_COUNTER_THROUGHPUT_INTEL                         int
	PERFQUERY_COUNTER_TIMESTAMP_INTEL                          int
	PERFQUERY_DONOT_FLUSH_INTEL                                int
	PERFQUERY_FLUSH_INTEL                                      int
	PERFQUERY_GLOBAL_CONTEXT_INTEL                             int
	PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL                      int
	PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL                      int
	PERFQUERY_SINGLE_CONTEXT_INTEL                             int
	PERFQUERY_WAIT_INTEL                                       int
	PINLIGHT_NV                                                int
	PIXEL_BUFFER_BARRIER_BIT                                   int
	PIXEL_PACK_BUFFER                                          int
	PIXEL_PACK_BUFFER_ARB                                      int
	PIXEL_PACK_BUFFER_BINDING                                  int
	PIXEL_PACK_BUFFER_BINDING_ARB                              int
	PIXEL_UNPACK_BUFFER                                        int
	PIXEL_UNPACK_BUFFER_ARB                                    int
	PIXEL_UNPACK_BUFFER_BINDING                                int
	PIXEL_UNPACK_BUFFER_BINDING_ARB                            int
	PLUS_CLAMPED_ALPHA_NV                                      int
	PLUS_CLAMPED_NV                                            int
	PLUS_DARKER_NV                                             int
	PLUS_NV                                                    int
	POINT                                                      int
	POINTS                                                     int
	POINT_FADE_THRESHOLD_SIZE                                  int
	POINT_SIZE                                                 int
	POINT_SIZE_GRANULARITY                                     int
	POINT_SIZE_RANGE                                           int
	POINT_SPRITE_COORD_ORIGIN                                  int
	POLYGON_MODE                                               int
	POLYGON_OFFSET_CLAMP                                       int
	POLYGON_OFFSET_CLAMP_EXT                                   int
	POLYGON_OFFSET_COMMAND_NV                                  int
	POLYGON_OFFSET_FACTOR                                      int
	POLYGON_OFFSET_FILL                                        int
	POLYGON_OFFSET_LINE                                        int
	POLYGON_OFFSET_POINT                                       int
	POLYGON_OFFSET_UNITS                                       int
	POLYGON_SMOOTH                                             int
	POLYGON_SMOOTH_HINT                                        int
	PRIMITIVES_GENERATED                                       int
	PRIMITIVES_SUBMITTED                                       int
	PRIMITIVES_SUBMITTED_ARB                                   int
	PRIMITIVE_BOUNDING_BOX_ARB                                 int
	PRIMITIVE_RESTART                                          int
	PRIMITIVE_RESTART_FIXED_INDEX                              int
	PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED                    int
	PRIMITIVE_RESTART_INDEX                                    int
	PROGRAM                                                    int
	PROGRAMMABLE_SAMPLE_LOCATION_ARB                           int
	PROGRAMMABLE_SAMPLE_LOCATION_NV                            int
	PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_ARB                int
	PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_NV                 int
	PROGRAM_BINARY_FORMATS                                     int
	PROGRAM_BINARY_LENGTH                                      int
	PROGRAM_BINARY_RETRIEVABLE_HINT                            int
	PROGRAM_INPUT                                              int
	PROGRAM_KHR                                                int
	PROGRAM_MATRIX_EXT                                         int
	PROGRAM_MATRIX_STACK_DEPTH_EXT                             int
	PROGRAM_OBJECT_EXT                                         int
	PROGRAM_OUTPUT                                             int
	PROGRAM_PIPELINE                                           int
	PROGRAM_PIPELINE_BINDING                                   int
	PROGRAM_PIPELINE_BINDING_EXT                               int
	PROGRAM_PIPELINE_KHR                                       int
	PROGRAM_PIPELINE_OBJECT_EXT                                int
	PROGRAM_POINT_SIZE                                         int
	PROGRAM_POINT_SIZE_ARB                                     int
	PROGRAM_SEPARABLE                                          int
	PROGRAM_SEPARABLE_EXT                                      int
	PROVOKING_VERTEX                                           int
	PROXY_TEXTURE_1D                                           int
	PROXY_TEXTURE_1D_ARRAY                                     int
	PROXY_TEXTURE_2D                                           int
	PROXY_TEXTURE_2D_ARRAY                                     int
	PROXY_TEXTURE_2D_MULTISAMPLE                               int
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         int
	PROXY_TEXTURE_3D                                           int
	PROXY_TEXTURE_CUBE_MAP                                     int
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               int
	PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB                           int
	PROXY_TEXTURE_RECTANGLE                                    int
	QUADRATIC_CURVE_TO_NV                                      int
	QUADS                                                      int
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   int
	QUERY                                                      int
	QUERY_BUFFER                                               int
	QUERY_BUFFER_BARRIER_BIT                                   int
	QUERY_BUFFER_BINDING                                       int
	QUERY_BY_REGION_NO_WAIT                                    int
	QUERY_BY_REGION_NO_WAIT_INVERTED                           int
	QUERY_BY_REGION_NO_WAIT_NV                                 int
	QUERY_BY_REGION_WAIT                                       int
	QUERY_BY_REGION_WAIT_INVERTED                              int
	QUERY_BY_REGION_WAIT_NV                                    int
	QUERY_COUNTER_BITS                                         int
	QUERY_KHR                                                  int
	QUERY_NO_WAIT                                              int
	QUERY_NO_WAIT_INVERTED                                     int
	QUERY_NO_WAIT_NV                                           int
	QUERY_OBJECT_EXT                                           int
	QUERY_RESULT                                               int
	QUERY_RESULT_AVAILABLE                                     int
	QUERY_RESULT_NO_WAIT                                       int
	QUERY_TARGET                                               int
	QUERY_WAIT                                                 int
	QUERY_WAIT_INVERTED                                        int
	QUERY_WAIT_NV                                              int
	R11F_G11F_B10F                                             int
	R16                                                        int
	R16F                                                       int
	R16I                                                       int
	R16UI                                                      int
	R16_SNORM                                                  int
	R32F                                                       int
	R32I                                                       int
	R32UI                                                      int
	R3_G3_B2                                                   int
	R8                                                         int
	R8I                                                        int
	R8UI                                                       int
	R8_SNORM                                                   int
	RASTERIZER_DISCARD                                         int
	RASTER_FIXED_SAMPLE_LOCATIONS_EXT                          int
	RASTER_MULTISAMPLE_EXT                                     int
	RASTER_SAMPLES_EXT                                         int
	READ_BUFFER                                                int
	READ_FRAMEBUFFER                                           int
	READ_FRAMEBUFFER_BINDING                                   int
	READ_ONLY                                                  int
	READ_PIXELS                                                int
	READ_PIXELS_FORMAT                                         int
	READ_PIXELS_TYPE                                           int
	READ_WRITE                                                 int
	RECT_NV                                                    int
	RED                                                        int
	RED_INTEGER                                                int
	RED_NV                                                     int
	REFERENCED_BY_COMPUTE_SHADER                               int
	REFERENCED_BY_FRAGMENT_SHADER                              int
	REFERENCED_BY_GEOMETRY_SHADER                              int
	REFERENCED_BY_TESS_CONTROL_SHADER                          int
	REFERENCED_BY_TESS_EVALUATION_SHADER                       int
	REFERENCED_BY_VERTEX_SHADER                                int
	RELATIVE_ARC_TO_NV                                         int
	RELATIVE_CONIC_CURVE_TO_NV                                 int
	RELATIVE_CUBIC_CURVE_TO_NV                                 int
	RELATIVE_HORIZONTAL_LINE_TO_NV                             int
	RELATIVE_LARGE_CCW_ARC_TO_NV                               int
	RELATIVE_LARGE_CW_ARC_TO_NV                                int
	RELATIVE_LINE_TO_NV                                        int
	RELATIVE_MOVE_TO_NV                                        int
	RELATIVE_QUADRATIC_CURVE_TO_NV                             int
	RELATIVE_RECT_NV                                           int
	RELATIVE_ROUNDED_RECT2_NV                                  int
	RELATIVE_ROUNDED_RECT4_NV                                  int
	RELATIVE_ROUNDED_RECT8_NV                                  int
	RELATIVE_ROUNDED_RECT_NV                                   int
	RELATIVE_SMALL_CCW_ARC_TO_NV                               int
	RELATIVE_SMALL_CW_ARC_TO_NV                                int
	RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV                          int
	RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV                      int
	RELATIVE_VERTICAL_LINE_TO_NV                               int
	RENDERBUFFER                                               int
	RENDERBUFFER_ALPHA_SIZE                                    int
	RENDERBUFFER_BINDING                                       int
	RENDERBUFFER_BLUE_SIZE                                     int
	RENDERBUFFER_COLOR_SAMPLES_NV                              int
	RENDERBUFFER_COVERAGE_SAMPLES_NV                           int
	RENDERBUFFER_DEPTH_SIZE                                    int
	RENDERBUFFER_GREEN_SIZE                                    int
	RENDERBUFFER_HEIGHT                                        int
	RENDERBUFFER_INTERNAL_FORMAT                               int
	RENDERBUFFER_RED_SIZE                                      int
	RENDERBUFFER_SAMPLES                                       int
	RENDERBUFFER_STENCIL_SIZE                                  int
	RENDERBUFFER_WIDTH                                         int
	RENDERER                                                   int
	REPEAT                                                     int
	REPLACE                                                    int
	RESET_NOTIFICATION_STRATEGY                                int
	RESET_NOTIFICATION_STRATEGY_ARB                            int
	RESET_NOTIFICATION_STRATEGY_KHR                            int
	RESTART_PATH_NV                                            int
	RG                                                         int
	RG16                                                       int
	RG16F                                                      int
	RG16I                                                      int
	RG16UI                                                     int
	RG16_SNORM                                                 int
	RG32F                                                      int
	RG32I                                                      int
	RG32UI                                                     int
	RG8                                                        int
	RG8I                                                       int
	RG8UI                                                      int
	RG8_SNORM                                                  int
	RGB                                                        int
	RGB10                                                      int
	RGB10_A2                                                   int
	RGB10_A2UI                                                 int
	RGB12                                                      int
	RGB16                                                      int
	RGB16F                                                     int
	RGB16I                                                     int
	RGB16UI                                                    int
	RGB16_SNORM                                                int
	RGB32F                                                     int
	RGB32I                                                     int
	RGB32UI                                                    int
	RGB4                                                       int
	RGB5                                                       int
	RGB565                                                     int
	RGB5_A1                                                    int
	RGB8                                                       int
	RGB8I                                                      int
	RGB8UI                                                     int
	RGB8_SNORM                                                 int
	RGB9_E5                                                    int
	RGBA                                                       int
	RGBA12                                                     int
	RGBA16                                                     int
	RGBA16F                                                    int
	RGBA16I                                                    int
	RGBA16UI                                                   int
	RGBA16_SNORM                                               int
	RGBA2                                                      int
	RGBA32F                                                    int
	RGBA32I                                                    int
	RGBA32UI                                                   int
	RGBA4                                                      int
	RGBA8                                                      int
	RGBA8I                                                     int
	RGBA8UI                                                    int
	RGBA8_SNORM                                                int
	RGBA_INTEGER                                               int
	RGB_422_APPLE                                              int
	RGB_INTEGER                                                int
	RGB_RAW_422_APPLE                                          int
	RG_INTEGER                                                 int
	RIGHT                                                      int
	ROUNDED_RECT2_NV                                           int
	ROUNDED_RECT4_NV                                           int
	ROUNDED_RECT8_NV                                           int
	ROUNDED_RECT_NV                                            int
	ROUND_NV                                                   int
	SAMPLER                                                    int
	SAMPLER_1D                                                 int
	SAMPLER_1D_ARRAY                                           int
	SAMPLER_1D_ARRAY_SHADOW                                    int
	SAMPLER_1D_SHADOW                                          int
	SAMPLER_2D                                                 int
	SAMPLER_2D_ARRAY                                           int
	SAMPLER_2D_ARRAY_SHADOW                                    int
	SAMPLER_2D_MULTISAMPLE                                     int
	SAMPLER_2D_MULTISAMPLE_ARRAY                               int
	SAMPLER_2D_RECT                                            int
	SAMPLER_2D_RECT_SHADOW                                     int
	SAMPLER_2D_SHADOW                                          int
	SAMPLER_3D                                                 int
	SAMPLER_BINDING                                            int
	SAMPLER_BUFFER                                             int
	SAMPLER_CUBE                                               int
	SAMPLER_CUBE_MAP_ARRAY                                     int
	SAMPLER_CUBE_MAP_ARRAY_ARB                                 int
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              int
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB                          int
	SAMPLER_CUBE_SHADOW                                        int
	SAMPLER_KHR                                                int
	SAMPLES                                                    int
	SAMPLES_PASSED                                             int
	SAMPLE_ALPHA_TO_COVERAGE                                   int
	SAMPLE_ALPHA_TO_ONE                                        int
	SAMPLE_BUFFERS                                             int
	SAMPLE_COVERAGE                                            int
	SAMPLE_COVERAGE_INVERT                                     int
	SAMPLE_COVERAGE_VALUE                                      int
	SAMPLE_LOCATION_ARB                                        int
	SAMPLE_LOCATION_NV                                         int
	SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_ARB                      int
	SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_NV                       int
	SAMPLE_LOCATION_PIXEL_GRID_WIDTH_ARB                       int
	SAMPLE_LOCATION_PIXEL_GRID_WIDTH_NV                        int
	SAMPLE_LOCATION_SUBPIXEL_BITS_ARB                          int
	SAMPLE_LOCATION_SUBPIXEL_BITS_NV                           int
	SAMPLE_MASK                                                int
	SAMPLE_MASK_VALUE                                          int
	SAMPLE_POSITION                                            int
	SAMPLE_SHADING                                             int
	SAMPLE_SHADING_ARB                                         int
	SCISSOR_BOX                                                int
	SCISSOR_COMMAND_NV                                         int
	SCISSOR_TEST                                               int
	SCREEN_KHR                                                 int
	SCREEN_NV                                                  int
	SECONDARY_COLOR_ARRAY_ADDRESS_NV                           int
	SECONDARY_COLOR_ARRAY_LENGTH_NV                            int
	SEPARATE_ATTRIBS                                           int
	SET                                                        int
	SHADER                                                     int
	SHADER_BINARY_FORMATS                                      int
	SHADER_BINARY_FORMAT_SPIR_V                                int
	SHADER_BINARY_FORMAT_SPIR_V_ARB                            int
	SHADER_COMPILER                                            int
	SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV                        int
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            int
	SHADER_IMAGE_ATOMIC                                        int
	SHADER_IMAGE_LOAD                                          int
	SHADER_IMAGE_STORE                                         int
	SHADER_INCLUDE_ARB                                         int
	SHADER_KHR                                                 int
	SHADER_OBJECT_EXT                                          int
	SHADER_SOURCE_LENGTH                                       int
	SHADER_STORAGE_BARRIER_BIT                                 int
	SHADER_STORAGE_BLOCK                                       int
	SHADER_STORAGE_BUFFER                                      int
	SHADER_STORAGE_BUFFER_BINDING                              int
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     int
	SHADER_STORAGE_BUFFER_SIZE                                 int
	SHADER_STORAGE_BUFFER_START                                int
	SHADER_TYPE                                                int
	SHADING_LANGUAGE_VERSION                                   int
	SHARED_EDGE_NV                                             int
	SHORT                                                      int
	SIGNALED                                                   int
	SIGNED_NORMALIZED                                          int
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        int
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       int
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      int
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     int
	SKIP_DECODE_EXT                                            int
	SKIP_MISSING_GLYPH_NV                                      int
	SMALL_CCW_ARC_TO_NV                                        int
	SMALL_CW_ARC_TO_NV                                         int
	SMOOTH_CUBIC_CURVE_TO_NV                                   int
	SMOOTH_LINE_WIDTH_GRANULARITY                              int
	SMOOTH_LINE_WIDTH_RANGE                                    int
	SMOOTH_POINT_SIZE_GRANULARITY                              int
	SMOOTH_POINT_SIZE_RANGE                                    int
	SMOOTH_QUADRATIC_CURVE_TO_NV                               int
	SM_COUNT_NV                                                int
	SOFTLIGHT_KHR                                              int
	SOFTLIGHT_NV                                               int
	SPARSE_BUFFER_PAGE_SIZE_ARB                                int
	SPARSE_STORAGE_BIT_ARB                                     int
	SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB                 int
	SPIR_V_BINARY                                              int
	SPIR_V_BINARY_ARB                                          int
	SPIR_V_EXTENSIONS                                          int
	SQUARE_NV                                                  int
	SRC1_ALPHA                                                 int
	SRC1_COLOR                                                 int
	SRC_ALPHA                                                  int
	SRC_ALPHA_SATURATE                                         int
	SRC_ATOP_NV                                                int
	SRC_COLOR                                                  int
	SRC_IN_NV                                                  int
	SRC_NV                                                     int
	SRC_OUT_NV                                                 int
	SRC_OVER_NV                                                int
	SRGB                                                       int
	SRGB8                                                      int
	SRGB8_ALPHA8                                               int
	SRGB_ALPHA                                                 int
	SRGB_DECODE_ARB                                            int
	SRGB_READ                                                  int
	SRGB_WRITE                                                 int
	STACK_OVERFLOW                                             int
	STACK_OVERFLOW_KHR                                         int
	STACK_UNDERFLOW                                            int
	STACK_UNDERFLOW_KHR                                        int
	STANDARD_FONT_FORMAT_NV                                    int
	STANDARD_FONT_NAME_NV                                      int
	STATIC_COPY                                                int
	STATIC_DRAW                                                int
	STATIC_READ                                                int
	STENCIL                                                    int
	STENCIL_ATTACHMENT                                         int
	STENCIL_BACK_FAIL                                          int
	STENCIL_BACK_FUNC                                          int
	STENCIL_BACK_PASS_DEPTH_FAIL                               int
	STENCIL_BACK_PASS_DEPTH_PASS                               int
	STENCIL_BACK_REF                                           int
	STENCIL_BACK_VALUE_MASK                                    int
	STENCIL_BACK_WRITEMASK                                     int
	STENCIL_BUFFER_BIT                                         int
	STENCIL_CLEAR_VALUE                                        int
	STENCIL_COMPONENTS                                         int
	STENCIL_FAIL                                               int
	STENCIL_FUNC                                               int
	STENCIL_INDEX                                              int
	STENCIL_INDEX1                                             int
	STENCIL_INDEX16                                            int
	STENCIL_INDEX4                                             int
	STENCIL_INDEX8                                             int
	STENCIL_PASS_DEPTH_FAIL                                    int
	STENCIL_PASS_DEPTH_PASS                                    int
	STENCIL_REF                                                int
	STENCIL_REF_COMMAND_NV                                     int
	STENCIL_RENDERABLE                                         int
	STENCIL_SAMPLES_NV                                         int
	STENCIL_TEST                                               int
	STENCIL_VALUE_MASK                                         int
	STENCIL_WRITEMASK                                          int
	STEREO                                                     int
	STREAM_COPY                                                int
	STREAM_DRAW                                                int
	STREAM_READ                                                int
	SUBPIXEL_BITS                                              int
	SUBPIXEL_PRECISION_BIAS_X_BITS_NV                          int
	SUBPIXEL_PRECISION_BIAS_Y_BITS_NV                          int
	SUPERSAMPLE_SCALE_X_NV                                     int
	SUPERSAMPLE_SCALE_Y_NV                                     int
	SYNC_CL_EVENT_ARB                                          int
	SYNC_CL_EVENT_COMPLETE_ARB                                 int
	SYNC_CONDITION                                             int
	SYNC_FENCE                                                 int
	SYNC_FLAGS                                                 int
	SYNC_FLUSH_COMMANDS_BIT                                    int
	SYNC_GPU_COMMANDS_COMPLETE                                 int
	SYNC_STATUS                                                int
	SYSTEM_FONT_NAME_NV                                        int
	TERMINATE_SEQUENCE_COMMAND_NV                              int
	TESS_CONTROL_OUTPUT_VERTICES                               int
	TESS_CONTROL_SHADER                                        int
	TESS_CONTROL_SHADER_BIT                                    int
	TESS_CONTROL_SHADER_PATCHES                                int
	TESS_CONTROL_SHADER_PATCHES_ARB                            int
	TESS_CONTROL_SUBROUTINE                                    int
	TESS_CONTROL_SUBROUTINE_UNIFORM                            int
	TESS_CONTROL_TEXTURE                                       int
	TESS_EVALUATION_SHADER                                     int
	TESS_EVALUATION_SHADER_BIT                                 int
	TESS_EVALUATION_SHADER_INVOCATIONS                         int
	TESS_EVALUATION_SHADER_INVOCATIONS_ARB                     int
	TESS_EVALUATION_SUBROUTINE                                 int
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         int
	TESS_EVALUATION_TEXTURE                                    int
	TESS_GEN_MODE                                              int
	TESS_GEN_POINT_MODE                                        int
	TESS_GEN_SPACING                                           int
	TESS_GEN_VERTEX_ORDER                                      int
	TEXTURE                                                    int
	TEXTURE0                                                   int
	TEXTURE1                                                   int
	TEXTURE10                                                  int
	TEXTURE11                                                  int
	TEXTURE12                                                  int
	TEXTURE13                                                  int
	TEXTURE14                                                  int
	TEXTURE15                                                  int
	TEXTURE16                                                  int
	TEXTURE17                                                  int
	TEXTURE18                                                  int
	TEXTURE19                                                  int
	TEXTURE2                                                   int
	TEXTURE20                                                  int
	TEXTURE21                                                  int
	TEXTURE22                                                  int
	TEXTURE23                                                  int
	TEXTURE24                                                  int
	TEXTURE25                                                  int
	TEXTURE26                                                  int
	TEXTURE27                                                  int
	TEXTURE28                                                  int
	TEXTURE29                                                  int
	TEXTURE3                                                   int
	TEXTURE30                                                  int
	TEXTURE31                                                  int
	TEXTURE4                                                   int
	TEXTURE5                                                   int
	TEXTURE6                                                   int
	TEXTURE7                                                   int
	TEXTURE8                                                   int
	TEXTURE9                                                   int
	TEXTURE_1D                                                 int
	TEXTURE_1D_ARRAY                                           int
	TEXTURE_2D                                                 int
	TEXTURE_2D_ARRAY                                           int
	TEXTURE_2D_MULTISAMPLE                                     int
	TEXTURE_2D_MULTISAMPLE_ARRAY                               int
	TEXTURE_3D                                                 int
	TEXTURE_ALPHA_SIZE                                         int
	TEXTURE_ALPHA_TYPE                                         int
	TEXTURE_BASE_LEVEL                                         int
	TEXTURE_BINDING_1D                                         int
	TEXTURE_BINDING_1D_ARRAY                                   int
	TEXTURE_BINDING_2D                                         int
	TEXTURE_BINDING_2D_ARRAY                                   int
	TEXTURE_BINDING_2D_MULTISAMPLE                             int
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       int
	TEXTURE_BINDING_3D                                         int
	TEXTURE_BINDING_BUFFER                                     int
	TEXTURE_BINDING_BUFFER_ARB                                 int
	TEXTURE_BINDING_CUBE_MAP                                   int
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             int
	TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB                         int
	TEXTURE_BINDING_RECTANGLE                                  int
	TEXTURE_BLUE_SIZE                                          int
	TEXTURE_BLUE_TYPE                                          int
	TEXTURE_BORDER_COLOR                                       int
	TEXTURE_BUFFER                                             int
	TEXTURE_BUFFER_ARB                                         int
	TEXTURE_BUFFER_BINDING                                     int
	TEXTURE_BUFFER_DATA_STORE_BINDING                          int
	TEXTURE_BUFFER_DATA_STORE_BINDING_ARB                      int
	TEXTURE_BUFFER_FORMAT_ARB                                  int
	TEXTURE_BUFFER_OFFSET                                      int
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            int
	TEXTURE_BUFFER_SIZE                                        int
	TEXTURE_COMPARE_FUNC                                       int
	TEXTURE_COMPARE_MODE                                       int
	TEXTURE_COMPRESSED                                         int
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            int
	TEXTURE_COMPRESSED_BLOCK_SIZE                              int
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             int
	TEXTURE_COMPRESSED_IMAGE_SIZE                              int
	TEXTURE_COMPRESSION_HINT                                   int
	TEXTURE_COORD_ARRAY_ADDRESS_NV                             int
	TEXTURE_COORD_ARRAY_LENGTH_NV                              int
	TEXTURE_CUBE_MAP                                           int
	TEXTURE_CUBE_MAP_ARRAY                                     int
	TEXTURE_CUBE_MAP_ARRAY_ARB                                 int
	TEXTURE_CUBE_MAP_NEGATIVE_X                                int
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                int
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                int
	TEXTURE_CUBE_MAP_POSITIVE_X                                int
	TEXTURE_CUBE_MAP_POSITIVE_Y                                int
	TEXTURE_CUBE_MAP_POSITIVE_Z                                int
	TEXTURE_CUBE_MAP_SEAMLESS                                  int
	TEXTURE_DEPTH                                              int
	TEXTURE_DEPTH_SIZE                                         int
	TEXTURE_DEPTH_TYPE                                         int
	TEXTURE_FETCH_BARRIER_BIT                                  int
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             int
	TEXTURE_GATHER                                             int
	TEXTURE_GATHER_SHADOW                                      int
	TEXTURE_GREEN_SIZE                                         int
	TEXTURE_GREEN_TYPE                                         int
	TEXTURE_HEIGHT                                             int
	TEXTURE_IMAGE_FORMAT                                       int
	TEXTURE_IMAGE_TYPE                                         int
	TEXTURE_IMMUTABLE_FORMAT                                   int
	TEXTURE_IMMUTABLE_LEVELS                                   int
	TEXTURE_INTERNAL_FORMAT                                    int
	TEXTURE_LOD_BIAS                                           int
	TEXTURE_MAG_FILTER                                         int
	TEXTURE_MAX_ANISOTROPY                                     int
	TEXTURE_MAX_LEVEL                                          int
	TEXTURE_MAX_LOD                                            int
	TEXTURE_MIN_FILTER                                         int
	TEXTURE_MIN_LOD                                            int
	TEXTURE_RECTANGLE                                          int
	TEXTURE_REDUCTION_MODE_ARB                                 int
	TEXTURE_REDUCTION_MODE_EXT                                 int
	TEXTURE_RED_SIZE                                           int
	TEXTURE_RED_TYPE                                           int
	TEXTURE_SAMPLES                                            int
	TEXTURE_SHADOW                                             int
	TEXTURE_SHARED_SIZE                                        int
	TEXTURE_SPARSE_ARB                                         int
	TEXTURE_SRGB_DECODE_EXT                                    int
	TEXTURE_STENCIL_SIZE                                       int
	TEXTURE_SWIZZLE_A                                          int
	TEXTURE_SWIZZLE_B                                          int
	TEXTURE_SWIZZLE_G                                          int
	TEXTURE_SWIZZLE_R                                          int
	TEXTURE_SWIZZLE_RGBA                                       int
	TEXTURE_TARGET                                             int
	TEXTURE_UPDATE_BARRIER_BIT                                 int
	TEXTURE_VIEW                                               int
	TEXTURE_VIEW_MIN_LAYER                                     int
	TEXTURE_VIEW_MIN_LEVEL                                     int
	TEXTURE_VIEW_NUM_LAYERS                                    int
	TEXTURE_VIEW_NUM_LEVELS                                    int
	TEXTURE_WIDTH                                              int
	TEXTURE_WRAP_R                                             int
	TEXTURE_WRAP_S                                             int
	TEXTURE_WRAP_T                                             int
	TIMEOUT_EXPIRED                                            int
	TIMEOUT_IGNORED                                            uint64
	TIMESTAMP                                                  int
	TIME_ELAPSED                                               int
	TOP_LEVEL_ARRAY_SIZE                                       int
	TOP_LEVEL_ARRAY_STRIDE                                     int
	TRANSFORM_FEEDBACK                                         int
	TRANSFORM_FEEDBACK_ACTIVE                                  int
	TRANSFORM_FEEDBACK_BARRIER_BIT                             int
	TRANSFORM_FEEDBACK_BINDING                                 int
	TRANSFORM_FEEDBACK_BUFFER                                  int
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           int
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          int
	TRANSFORM_FEEDBACK_BUFFER_INDEX                            int
	TRANSFORM_FEEDBACK_BUFFER_MODE                             int
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           int
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             int
	TRANSFORM_FEEDBACK_BUFFER_START                            int
	TRANSFORM_FEEDBACK_BUFFER_STRIDE                           int
	TRANSFORM_FEEDBACK_OVERFLOW                                int
	TRANSFORM_FEEDBACK_OVERFLOW_ARB                            int
	TRANSFORM_FEEDBACK_PAUSED                                  int
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      int
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW                         int
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB                     int
	TRANSFORM_FEEDBACK_VARYING                                 int
	TRANSFORM_FEEDBACK_VARYINGS                                int
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      int
	TRANSLATE_2D_NV                                            int
	TRANSLATE_3D_NV                                            int
	TRANSLATE_X_NV                                             int
	TRANSLATE_Y_NV                                             int
	TRANSPOSE_AFFINE_2D_NV                                     int
	TRANSPOSE_AFFINE_3D_NV                                     int
	TRANSPOSE_PROGRAM_MATRIX_EXT                               int
	TRIANGLES                                                  int
	TRIANGLES_ADJACENCY                                        int
	TRIANGLES_ADJACENCY_ARB                                    int
	TRIANGLE_FAN                                               int
	TRIANGLE_STRIP                                             int
	TRIANGLE_STRIP_ADJACENCY                                   int
	TRIANGLE_STRIP_ADJACENCY_ARB                               int
	TRIANGULAR_NV                                              int
	TRUE                                                       int
	TYPE                                                       int
	UNCORRELATED_NV                                            int
	UNDEFINED_VERTEX                                           int
	UNIFORM                                                    int
	UNIFORM_ADDRESS_COMMAND_NV                                 int
	UNIFORM_ARRAY_STRIDE                                       int
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        int
	UNIFORM_BARRIER_BIT                                        int
	UNIFORM_BLOCK                                              int
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              int
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       int
	UNIFORM_BLOCK_BINDING                                      int
	UNIFORM_BLOCK_DATA_SIZE                                    int
	UNIFORM_BLOCK_INDEX                                        int
	UNIFORM_BLOCK_NAME_LENGTH                                  int
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 int
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                int
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                int
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            int
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         int
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  int
	UNIFORM_BUFFER                                             int
	UNIFORM_BUFFER_ADDRESS_NV                                  int
	UNIFORM_BUFFER_BINDING                                     int
	UNIFORM_BUFFER_LENGTH_NV                                   int
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            int
	UNIFORM_BUFFER_SIZE                                        int
	UNIFORM_BUFFER_START                                       int
	UNIFORM_BUFFER_UNIFIED_NV                                  int
	UNIFORM_IS_ROW_MAJOR                                       int
	UNIFORM_MATRIX_STRIDE                                      int
	UNIFORM_NAME_LENGTH                                        int
	UNIFORM_OFFSET                                             int
	UNIFORM_SIZE                                               int
	UNIFORM_TYPE                                               int
	UNKNOWN_CONTEXT_RESET                                      int
	UNKNOWN_CONTEXT_RESET_ARB                                  int
	UNKNOWN_CONTEXT_RESET_KHR                                  int
	UNPACK_ALIGNMENT                                           int
	UNPACK_COMPRESSED_BLOCK_DEPTH                              int
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             int
	UNPACK_COMPRESSED_BLOCK_SIZE                               int
	UNPACK_COMPRESSED_BLOCK_WIDTH                              int
	UNPACK_IMAGE_HEIGHT                                        int
	UNPACK_LSB_FIRST                                           int
	UNPACK_ROW_LENGTH                                          int
	UNPACK_SKIP_IMAGES                                         int
	UNPACK_SKIP_PIXELS                                         int
	UNPACK_SKIP_ROWS                                           int
	UNPACK_SWAP_BYTES                                          int
	UNSIGNALED                                                 int
	UNSIGNED_BYTE                                              int
	UNSIGNED_BYTE_2_3_3_REV                                    int
	UNSIGNED_BYTE_3_3_2                                        int
	UNSIGNED_INT                                               int
	UNSIGNED_INT16_NV                                          int
	UNSIGNED_INT16_VEC2_NV                                     int
	UNSIGNED_INT16_VEC3_NV                                     int
	UNSIGNED_INT16_VEC4_NV                                     int
	UNSIGNED_INT64_AMD                                         int
	UNSIGNED_INT64_ARB                                         int
	UNSIGNED_INT64_NV                                          int
	UNSIGNED_INT64_VEC2_ARB                                    int
	UNSIGNED_INT64_VEC2_NV                                     int
	UNSIGNED_INT64_VEC3_ARB                                    int
	UNSIGNED_INT64_VEC3_NV                                     int
	UNSIGNED_INT64_VEC4_ARB                                    int
	UNSIGNED_INT64_VEC4_NV                                     int
	UNSIGNED_INT8_NV                                           int
	UNSIGNED_INT8_VEC2_NV                                      int
	UNSIGNED_INT8_VEC3_NV                                      int
	UNSIGNED_INT8_VEC4_NV                                      int
	UNSIGNED_INT_10F_11F_11F_REV                               int
	UNSIGNED_INT_10_10_10_2                                    int
	UNSIGNED_INT_24_8                                          int
	UNSIGNED_INT_2_10_10_10_REV                                int
	UNSIGNED_INT_5_9_9_9_REV                                   int
	UNSIGNED_INT_8_8_8_8                                       int
	UNSIGNED_INT_8_8_8_8_REV                                   int
	UNSIGNED_INT_ATOMIC_COUNTER                                int
	UNSIGNED_INT_IMAGE_1D                                      int
	UNSIGNED_INT_IMAGE_1D_ARRAY                                int
	UNSIGNED_INT_IMAGE_2D                                      int
	UNSIGNED_INT_IMAGE_2D_ARRAY                                int
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          int
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    int
	UNSIGNED_INT_IMAGE_2D_RECT                                 int
	UNSIGNED_INT_IMAGE_3D                                      int
	UNSIGNED_INT_IMAGE_BUFFER                                  int
	UNSIGNED_INT_IMAGE_CUBE                                    int
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          int
	UNSIGNED_INT_SAMPLER_1D                                    int
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              int
	UNSIGNED_INT_SAMPLER_2D                                    int
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              int
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        int
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  int
	UNSIGNED_INT_SAMPLER_2D_RECT                               int
	UNSIGNED_INT_SAMPLER_3D                                    int
	UNSIGNED_INT_SAMPLER_BUFFER                                int
	UNSIGNED_INT_SAMPLER_CUBE                                  int
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        int
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB                    int
	UNSIGNED_INT_VEC2                                          int
	UNSIGNED_INT_VEC3                                          int
	UNSIGNED_INT_VEC4                                          int
	UNSIGNED_NORMALIZED                                        int
	UNSIGNED_SHORT                                             int
	UNSIGNED_SHORT_1_5_5_5_REV                                 int
	UNSIGNED_SHORT_4_4_4_4                                     int
	UNSIGNED_SHORT_4_4_4_4_REV                                 int
	UNSIGNED_SHORT_5_5_5_1                                     int
	UNSIGNED_SHORT_5_6_5                                       int
	UNSIGNED_SHORT_5_6_5_REV                                   int
	UNSIGNED_SHORT_8_8_APPLE                                   int
	UNSIGNED_SHORT_8_8_REV_APPLE                               int
	UPPER_LEFT                                                 int
	USE_MISSING_GLYPH_NV                                       int
	UTF16_NV                                                   int
	UTF8_NV                                                    int
	VALIDATE_STATUS                                            int
	VENDOR                                                     int
	VERSION                                                    int
	VERTEX_ARRAY                                               int
	VERTEX_ARRAY_ADDRESS_NV                                    int
	VERTEX_ARRAY_BINDING                                       int
	VERTEX_ARRAY_KHR                                           int
	VERTEX_ARRAY_LENGTH_NV                                     int
	VERTEX_ARRAY_OBJECT_EXT                                    int
	VERTEX_ATTRIB_ARRAY_ADDRESS_NV                             int
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            int
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         int
	VERTEX_ATTRIB_ARRAY_DIVISOR                                int
	VERTEX_ATTRIB_ARRAY_DIVISOR_ARB                            int
	VERTEX_ATTRIB_ARRAY_ENABLED                                int
	VERTEX_ATTRIB_ARRAY_INTEGER                                int
	VERTEX_ATTRIB_ARRAY_LENGTH_NV                              int
	VERTEX_ATTRIB_ARRAY_LONG                                   int
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             int
	VERTEX_ATTRIB_ARRAY_POINTER                                int
	VERTEX_ATTRIB_ARRAY_SIZE                                   int
	VERTEX_ATTRIB_ARRAY_STRIDE                                 int
	VERTEX_ATTRIB_ARRAY_TYPE                                   int
	VERTEX_ATTRIB_ARRAY_UNIFIED_NV                             int
	VERTEX_ATTRIB_BINDING                                      int
	VERTEX_ATTRIB_RELATIVE_OFFSET                              int
	VERTEX_BINDING_BUFFER                                      int
	VERTEX_BINDING_DIVISOR                                     int
	VERTEX_BINDING_OFFSET                                      int
	VERTEX_BINDING_STRIDE                                      int
	VERTEX_PROGRAM_POINT_SIZE                                  int
	VERTEX_SHADER                                              int
	VERTEX_SHADER_BIT                                          int
	VERTEX_SHADER_BIT_EXT                                      int
	VERTEX_SHADER_INVOCATIONS                                  int
	VERTEX_SHADER_INVOCATIONS_ARB                              int
	VERTEX_SUBROUTINE                                          int
	VERTEX_SUBROUTINE_UNIFORM                                  int
	VERTEX_TEXTURE                                             int
	VERTICAL_LINE_TO_NV                                        int
	VERTICES_SUBMITTED                                         int
	VERTICES_SUBMITTED_ARB                                     int
	VIEWPORT                                                   int
	VIEWPORT_BOUNDS_RANGE                                      int
	VIEWPORT_COMMAND_NV                                        int
	VIEWPORT_INDEX_PROVOKING_VERTEX                            int
	VIEWPORT_POSITION_W_SCALE_NV                               int
	VIEWPORT_POSITION_W_SCALE_X_COEFF_NV                       int
	VIEWPORT_POSITION_W_SCALE_Y_COEFF_NV                       int
	VIEWPORT_SUBPIXEL_BITS                                     int
	VIEWPORT_SWIZZLE_NEGATIVE_W_NV                             int
	VIEWPORT_SWIZZLE_NEGATIVE_X_NV                             int
	VIEWPORT_SWIZZLE_NEGATIVE_Y_NV                             int
	VIEWPORT_SWIZZLE_NEGATIVE_Z_NV                             int
	VIEWPORT_SWIZZLE_POSITIVE_W_NV                             int
	VIEWPORT_SWIZZLE_POSITIVE_X_NV                             int
	VIEWPORT_SWIZZLE_POSITIVE_Y_NV                             int
	VIEWPORT_SWIZZLE_POSITIVE_Z_NV                             int
	VIEWPORT_SWIZZLE_W_NV                                      int
	VIEWPORT_SWIZZLE_X_NV                                      int
	VIEWPORT_SWIZZLE_Y_NV                                      int
	VIEWPORT_SWIZZLE_Z_NV                                      int
	VIEW_CLASS_128_BITS                                        int
	VIEW_CLASS_16_BITS                                         int
	VIEW_CLASS_24_BITS                                         int
	VIEW_CLASS_32_BITS                                         int
	VIEW_CLASS_48_BITS                                         int
	VIEW_CLASS_64_BITS                                         int
	VIEW_CLASS_8_BITS                                          int
	VIEW_CLASS_96_BITS                                         int
	VIEW_CLASS_BPTC_FLOAT                                      int
	VIEW_CLASS_BPTC_UNORM                                      int
	VIEW_CLASS_RGTC1_RED                                       int
	VIEW_CLASS_RGTC2_RG                                        int
	VIEW_CLASS_S3TC_DXT1_RGB                                   int
	VIEW_CLASS_S3TC_DXT1_RGBA                                  int
	VIEW_CLASS_S3TC_DXT3_RGBA                                  int
	VIEW_CLASS_S3TC_DXT5_RGBA                                  int
	VIEW_COMPATIBILITY_CLASS                                   int
	VIRTUAL_PAGE_SIZE_INDEX_ARB                                int
	VIRTUAL_PAGE_SIZE_X_ARB                                    int
	VIRTUAL_PAGE_SIZE_Y_ARB                                    int
	VIRTUAL_PAGE_SIZE_Z_ARB                                    int
	VIVIDLIGHT_NV                                              int
	WAIT_FAILED                                                int
	WARPS_PER_SM_NV                                            int
	WARP_SIZE_NV                                               int
	WEIGHTED_AVERAGE_ARB                                       int
	WEIGHTED_AVERAGE_EXT                                       int
	WINDOW_RECTANGLE_EXT                                       int
	WINDOW_RECTANGLE_MODE_EXT                                  int
	WRITE_ONLY                                                 int
	XOR                                                        int
	XOR_NV                                                     int
	ZERO                                                       int
	ZERO_TO_ONE                                                int
}

func NewContext() *Context {
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}
	return &Context{
		ACCUM_ADJACENT_PAIRS_NV:              gl.ACCUM_ADJACENT_PAIRS_NV,
		ACTIVE_ATOMIC_COUNTER_BUFFERS:        gl.ACTIVE_ATOMIC_COUNTER_BUFFERS,
		ACTIVE_ATTRIBUTES:                    gl.ACTIVE_ATTRIBUTES,
		ACTIVE_ATTRIBUTE_MAX_LENGTH:          gl.ACTIVE_ATTRIBUTE_MAX_LENGTH,
		ACTIVE_PROGRAM:                       gl.ACTIVE_PROGRAM,
		ACTIVE_PROGRAM_EXT:                   gl.ACTIVE_PROGRAM_EXT,
		ACTIVE_RESOURCES:                     gl.ACTIVE_RESOURCES,
		ACTIVE_SUBROUTINES:                   gl.ACTIVE_SUBROUTINES,
		ACTIVE_SUBROUTINE_MAX_LENGTH:         gl.ACTIVE_SUBROUTINE_MAX_LENGTH,
		ACTIVE_SUBROUTINE_UNIFORMS:           gl.ACTIVE_SUBROUTINE_UNIFORMS,
		ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS:  gl.ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS,
		ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH: gl.ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH,
		ACTIVE_TEXTURE:                       gl.ACTIVE_TEXTURE,
		ACTIVE_UNIFORMS:                      gl.ACTIVE_UNIFORMS,
		ACTIVE_UNIFORM_BLOCKS:                gl.ACTIVE_UNIFORM_BLOCKS,
		ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH: gl.ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH,
		ACTIVE_UNIFORM_MAX_LENGTH:            gl.ACTIVE_UNIFORM_MAX_LENGTH,
		ACTIVE_VARIABLES:                     gl.ACTIVE_VARIABLES,
		ADJACENT_PAIRS_NV:                    gl.ADJACENT_PAIRS_NV,
		AFFINE_2D_NV:                         gl.AFFINE_2D_NV,
		AFFINE_3D_NV:                         gl.AFFINE_3D_NV,
		ALIASED_LINE_WIDTH_RANGE:             gl.ALIASED_LINE_WIDTH_RANGE,
		ALL_BARRIER_BITS:                     gl.ALL_BARRIER_BITS,
		ALL_SHADER_BITS:                      gl.ALL_SHADER_BITS,
		ALL_SHADER_BITS_EXT:                  gl.ALL_SHADER_BITS_EXT,
		ALPHA:                                gl.ALPHA,
		ALPHA_REF_COMMAND_NV:                 gl.ALPHA_REF_COMMAND_NV,
		ALREADY_SIGNALED:                     gl.ALREADY_SIGNALED,
		ALWAYS:                               gl.ALWAYS,
		AND:                                  gl.AND,
		AND_INVERTED:                         gl.AND_INVERTED,
		AND_REVERSE:                          gl.AND_REVERSE,
		ANY_SAMPLES_PASSED:                   gl.ANY_SAMPLES_PASSED,
		ANY_SAMPLES_PASSED_CONSERVATIVE:      gl.ANY_SAMPLES_PASSED_CONSERVATIVE,
		ARC_TO_NV:                            gl.ARC_TO_NV,
		ARRAY_BUFFER:                         gl.ARRAY_BUFFER,
		ARRAY_BUFFER_BINDING:                 gl.ARRAY_BUFFER_BINDING,
		ARRAY_SIZE:                           gl.ARRAY_SIZE,
		ARRAY_STRIDE:                         gl.ARRAY_STRIDE,
		ATOMIC_COUNTER_BARRIER_BIT:           gl.ATOMIC_COUNTER_BARRIER_BIT,
		ATOMIC_COUNTER_BUFFER:                gl.ATOMIC_COUNTER_BUFFER,
		ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS:               gl.ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS,
		ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES:        gl.ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES,
		ATOMIC_COUNTER_BUFFER_BINDING:                              gl.ATOMIC_COUNTER_BUFFER_BINDING,
		ATOMIC_COUNTER_BUFFER_DATA_SIZE:                            gl.ATOMIC_COUNTER_BUFFER_DATA_SIZE,
		ATOMIC_COUNTER_BUFFER_INDEX:                                gl.ATOMIC_COUNTER_BUFFER_INDEX,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER:         gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER:        gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER:        gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER:    gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER: gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER,
		ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER:          gl.ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER,
		ATOMIC_COUNTER_BUFFER_SIZE:                                 gl.ATOMIC_COUNTER_BUFFER_SIZE,
		ATOMIC_COUNTER_BUFFER_START:                                gl.ATOMIC_COUNTER_BUFFER_START,
		ATTACHED_SHADERS:                                           gl.ATTACHED_SHADERS,
		ATTRIBUTE_ADDRESS_COMMAND_NV:                               gl.ATTRIBUTE_ADDRESS_COMMAND_NV,
		AUTO_GENERATE_MIPMAP:                                       gl.AUTO_GENERATE_MIPMAP,
		BACK:                                                       gl.BACK,
		BACK_LEFT:                                                  gl.BACK_LEFT,
		BACK_RIGHT:                                                 gl.BACK_RIGHT,
		BEVEL_NV:                                                   gl.BEVEL_NV,
		BGR:                                                        gl.BGR,
		BGRA:                                                       gl.BGRA,
		BGRA_INTEGER:                                               gl.BGRA_INTEGER,
		BGR_INTEGER:                                                gl.BGR_INTEGER,
		BLACKHOLE_RENDER_INTEL:                                     gl.BLACKHOLE_RENDER_INTEL,
		BLEND:                                                      gl.BLEND,
		BLEND_ADVANCED_COHERENT_KHR:                                gl.BLEND_ADVANCED_COHERENT_KHR,
		BLEND_ADVANCED_COHERENT_NV:                                 gl.BLEND_ADVANCED_COHERENT_NV,
		BLEND_COLOR:                                                gl.BLEND_COLOR,
		BLEND_COLOR_COMMAND_NV:                                     gl.BLEND_COLOR_COMMAND_NV,
		BLEND_DST:                                                  gl.BLEND_DST,
		BLEND_DST_ALPHA:                                            gl.BLEND_DST_ALPHA,
		BLEND_DST_RGB:                                              gl.BLEND_DST_RGB,
		BLEND_EQUATION:                                             gl.BLEND_EQUATION,
		BLEND_EQUATION_ALPHA:                                       gl.BLEND_EQUATION_ALPHA,
		BLEND_EQUATION_RGB:                                         gl.BLEND_EQUATION_RGB,
		BLEND_OVERLAP_NV:                                           gl.BLEND_OVERLAP_NV,
		BLEND_PREMULTIPLIED_SRC_NV:                                 gl.BLEND_PREMULTIPLIED_SRC_NV,
		BLEND_SRC:                                                  gl.BLEND_SRC,
		BLEND_SRC_ALPHA:                                            gl.BLEND_SRC_ALPHA,
		BLEND_SRC_RGB:                                              gl.BLEND_SRC_RGB,
		BLOCK_INDEX:                                                gl.BLOCK_INDEX,
		BLUE:                                                       gl.BLUE,
		BLUE_INTEGER:                                               gl.BLUE_INTEGER,
		BLUE_NV:                                                    gl.BLUE_NV,
		BOLD_BIT_NV:                                                gl.BOLD_BIT_NV,
		BOOL:                                                       gl.BOOL,
		BOOL_VEC2:                                                  gl.BOOL_VEC2,
		BOOL_VEC3:                                                  gl.BOOL_VEC3,
		BOOL_VEC4:                                                  gl.BOOL_VEC4,
		BOUNDING_BOX_NV:                                            gl.BOUNDING_BOX_NV,
		BOUNDING_BOX_OF_BOUNDING_BOXES_NV:                          gl.BOUNDING_BOX_OF_BOUNDING_BOXES_NV,
		BUFFER:                                                     gl.BUFFER,
		BUFFER_ACCESS:                                              gl.BUFFER_ACCESS,
		BUFFER_ACCESS_FLAGS:                                        gl.BUFFER_ACCESS_FLAGS,
		BUFFER_BINDING:                                             gl.BUFFER_BINDING,
		BUFFER_DATA_SIZE:                                           gl.BUFFER_DATA_SIZE,
		BUFFER_GPU_ADDRESS_NV:                                      gl.BUFFER_GPU_ADDRESS_NV,
		BUFFER_IMMUTABLE_STORAGE:                                   gl.BUFFER_IMMUTABLE_STORAGE,
		BUFFER_KHR:                                                 gl.BUFFER_KHR,
		BUFFER_MAPPED:                                              gl.BUFFER_MAPPED,
		BUFFER_MAP_LENGTH:                                          gl.BUFFER_MAP_LENGTH,
		BUFFER_MAP_OFFSET:                                          gl.BUFFER_MAP_OFFSET,
		BUFFER_MAP_POINTER:                                         gl.BUFFER_MAP_POINTER,
		BUFFER_OBJECT_EXT:                                          gl.BUFFER_OBJECT_EXT,
		BUFFER_SIZE:                                                gl.BUFFER_SIZE,
		BUFFER_STORAGE_FLAGS:                                       gl.BUFFER_STORAGE_FLAGS,
		BUFFER_UPDATE_BARRIER_BIT:                                  gl.BUFFER_UPDATE_BARRIER_BIT,
		BUFFER_USAGE:                                               gl.BUFFER_USAGE,
		BUFFER_VARIABLE:                                            gl.BUFFER_VARIABLE,
		BYTE:                                                       gl.BYTE,
		CAVEAT_SUPPORT:                                             gl.CAVEAT_SUPPORT,
		CCW:                                                        gl.CCW,
		CIRCULAR_CCW_ARC_TO_NV:                                     gl.CIRCULAR_CCW_ARC_TO_NV,
		CIRCULAR_CW_ARC_TO_NV:                                      gl.CIRCULAR_CW_ARC_TO_NV,
		CIRCULAR_TANGENT_ARC_TO_NV:                                 gl.CIRCULAR_TANGENT_ARC_TO_NV,
		CLAMP_READ_COLOR:                                           gl.CLAMP_READ_COLOR,
		CLAMP_TO_BORDER:                                            gl.CLAMP_TO_BORDER,
		CLAMP_TO_BORDER_ARB:                                        gl.CLAMP_TO_BORDER_ARB,
		CLAMP_TO_EDGE:                                              gl.CLAMP_TO_EDGE,
		CLEAR:                                                      gl.CLEAR,
		CLEAR_BUFFER:                                               gl.CLEAR_BUFFER,
		CLEAR_TEXTURE:                                              gl.CLEAR_TEXTURE,
		CLIENT_MAPPED_BUFFER_BARRIER_BIT:                           gl.CLIENT_MAPPED_BUFFER_BARRIER_BIT,
		CLIENT_STORAGE_BIT:                                         gl.CLIENT_STORAGE_BIT,
		CLIPPING_INPUT_PRIMITIVES:                                  gl.CLIPPING_INPUT_PRIMITIVES,
		CLIPPING_INPUT_PRIMITIVES_ARB:                              gl.CLIPPING_INPUT_PRIMITIVES_ARB,
		CLIPPING_OUTPUT_PRIMITIVES:                                 gl.CLIPPING_OUTPUT_PRIMITIVES,
		CLIPPING_OUTPUT_PRIMITIVES_ARB:                             gl.CLIPPING_OUTPUT_PRIMITIVES_ARB,
		CLIP_DEPTH_MODE:                                            gl.CLIP_DEPTH_MODE,
		CLIP_DISTANCE0:                                             gl.CLIP_DISTANCE0,
		CLIP_DISTANCE1:                                             gl.CLIP_DISTANCE1,
		CLIP_DISTANCE2:                                             gl.CLIP_DISTANCE2,
		CLIP_DISTANCE3:                                             gl.CLIP_DISTANCE3,
		CLIP_DISTANCE4:                                             gl.CLIP_DISTANCE4,
		CLIP_DISTANCE5:                                             gl.CLIP_DISTANCE5,
		CLIP_DISTANCE6:                                             gl.CLIP_DISTANCE6,
		CLIP_DISTANCE7:                                             gl.CLIP_DISTANCE7,
		CLIP_ORIGIN:                                                gl.CLIP_ORIGIN,
		CLOSE_PATH_NV:                                              gl.CLOSE_PATH_NV,
		COLOR:                                                      gl.COLOR,
		COLORBURN_KHR:                                              gl.COLORBURN_KHR,
		COLORBURN_NV:                                               gl.COLORBURN_NV,
		COLORDODGE_KHR:                                             gl.COLORDODGE_KHR,
		COLORDODGE_NV:                                              gl.COLORDODGE_NV,
		COLOR_ARRAY_ADDRESS_NV:                                     gl.COLOR_ARRAY_ADDRESS_NV,
		COLOR_ARRAY_LENGTH_NV:                                      gl.COLOR_ARRAY_LENGTH_NV,
		COLOR_ATTACHMENT0:                                          gl.COLOR_ATTACHMENT0,
		COLOR_ATTACHMENT1:                                          gl.COLOR_ATTACHMENT1,
		COLOR_ATTACHMENT10:                                         gl.COLOR_ATTACHMENT10,
		COLOR_ATTACHMENT11:                                         gl.COLOR_ATTACHMENT11,
		COLOR_ATTACHMENT12:                                         gl.COLOR_ATTACHMENT12,
		COLOR_ATTACHMENT13:                                         gl.COLOR_ATTACHMENT13,
		COLOR_ATTACHMENT14:                                         gl.COLOR_ATTACHMENT14,
		COLOR_ATTACHMENT15:                                         gl.COLOR_ATTACHMENT15,
		COLOR_ATTACHMENT16:                                         gl.COLOR_ATTACHMENT16,
		COLOR_ATTACHMENT17:                                         gl.COLOR_ATTACHMENT17,
		COLOR_ATTACHMENT18:                                         gl.COLOR_ATTACHMENT18,
		COLOR_ATTACHMENT19:                                         gl.COLOR_ATTACHMENT19,
		COLOR_ATTACHMENT2:                                          gl.COLOR_ATTACHMENT2,
		COLOR_ATTACHMENT20:                                         gl.COLOR_ATTACHMENT20,
		COLOR_ATTACHMENT21:                                         gl.COLOR_ATTACHMENT21,
		COLOR_ATTACHMENT22:                                         gl.COLOR_ATTACHMENT22,
		COLOR_ATTACHMENT23:                                         gl.COLOR_ATTACHMENT23,
		COLOR_ATTACHMENT24:                                         gl.COLOR_ATTACHMENT24,
		COLOR_ATTACHMENT25:                                         gl.COLOR_ATTACHMENT25,
		COLOR_ATTACHMENT26:                                         gl.COLOR_ATTACHMENT26,
		COLOR_ATTACHMENT27:                                         gl.COLOR_ATTACHMENT27,
		COLOR_ATTACHMENT28:                                         gl.COLOR_ATTACHMENT28,
		COLOR_ATTACHMENT29:                                         gl.COLOR_ATTACHMENT29,
		COLOR_ATTACHMENT3:                                          gl.COLOR_ATTACHMENT3,
		COLOR_ATTACHMENT30:                                         gl.COLOR_ATTACHMENT30,
		COLOR_ATTACHMENT31:                                         gl.COLOR_ATTACHMENT31,
		COLOR_ATTACHMENT4:                                          gl.COLOR_ATTACHMENT4,
		COLOR_ATTACHMENT5:                                          gl.COLOR_ATTACHMENT5,
		COLOR_ATTACHMENT6:                                          gl.COLOR_ATTACHMENT6,
		COLOR_ATTACHMENT7:                                          gl.COLOR_ATTACHMENT7,
		COLOR_ATTACHMENT8:                                          gl.COLOR_ATTACHMENT8,
		COLOR_ATTACHMENT9:                                          gl.COLOR_ATTACHMENT9,
		COLOR_BUFFER_BIT:                                           gl.COLOR_BUFFER_BIT,
		COLOR_CLEAR_VALUE:                                          gl.COLOR_CLEAR_VALUE,
		COLOR_COMPONENTS:                                           gl.COLOR_COMPONENTS,
		COLOR_ENCODING:                                             gl.COLOR_ENCODING,
		COLOR_LOGIC_OP:                                             gl.COLOR_LOGIC_OP,
		COLOR_RENDERABLE:                                           gl.COLOR_RENDERABLE,
		COLOR_SAMPLES_NV:                                           gl.COLOR_SAMPLES_NV,
		COLOR_WRITEMASK:                                            gl.COLOR_WRITEMASK,
		COMMAND_BARRIER_BIT:                                        gl.COMMAND_BARRIER_BIT,
		COMPARE_REF_TO_TEXTURE:                                     gl.COMPARE_REF_TO_TEXTURE,
		COMPATIBLE_SUBROUTINES:                                     gl.COMPATIBLE_SUBROUTINES,
		COMPILE_STATUS:                                             gl.COMPILE_STATUS,
		COMPLETION_STATUS_ARB:                                      gl.COMPLETION_STATUS_ARB,
		COMPLETION_STATUS_KHR:                                      gl.COMPLETION_STATUS_KHR,
		COMPRESSED_R11_EAC:                                         gl.COMPRESSED_R11_EAC,
		COMPRESSED_RED:                                             gl.COMPRESSED_RED,
		COMPRESSED_RED_RGTC1:                                       gl.COMPRESSED_RED_RGTC1,
		COMPRESSED_RG:                                              gl.COMPRESSED_RG,
		COMPRESSED_RG11_EAC:                                        gl.COMPRESSED_RG11_EAC,
		COMPRESSED_RGB:                                             gl.COMPRESSED_RGB,
		COMPRESSED_RGB8_ETC2:                                       gl.COMPRESSED_RGB8_ETC2,
		COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2:                   gl.COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2,
		COMPRESSED_RGBA:                                            gl.COMPRESSED_RGBA,
		COMPRESSED_RGBA8_ETC2_EAC:                                  gl.COMPRESSED_RGBA8_ETC2_EAC,
		COMPRESSED_RGBA_ASTC_10x10_KHR:                             gl.COMPRESSED_RGBA_ASTC_10x10_KHR,
		COMPRESSED_RGBA_ASTC_10x5_KHR:                              gl.COMPRESSED_RGBA_ASTC_10x5_KHR,
		COMPRESSED_RGBA_ASTC_10x6_KHR:                              gl.COMPRESSED_RGBA_ASTC_10x6_KHR,
		COMPRESSED_RGBA_ASTC_10x8_KHR:                              gl.COMPRESSED_RGBA_ASTC_10x8_KHR,
		COMPRESSED_RGBA_ASTC_12x10_KHR:                             gl.COMPRESSED_RGBA_ASTC_12x10_KHR,
		COMPRESSED_RGBA_ASTC_12x12_KHR:                             gl.COMPRESSED_RGBA_ASTC_12x12_KHR,
		COMPRESSED_RGBA_ASTC_4x4_KHR:                               gl.COMPRESSED_RGBA_ASTC_4x4_KHR,
		COMPRESSED_RGBA_ASTC_5x4_KHR:                               gl.COMPRESSED_RGBA_ASTC_5x4_KHR,
		COMPRESSED_RGBA_ASTC_5x5_KHR:                               gl.COMPRESSED_RGBA_ASTC_5x5_KHR,
		COMPRESSED_RGBA_ASTC_6x5_KHR:                               gl.COMPRESSED_RGBA_ASTC_6x5_KHR,
		COMPRESSED_RGBA_ASTC_6x6_KHR:                               gl.COMPRESSED_RGBA_ASTC_6x6_KHR,
		COMPRESSED_RGBA_ASTC_8x5_KHR:                               gl.COMPRESSED_RGBA_ASTC_8x5_KHR,
		COMPRESSED_RGBA_ASTC_8x6_KHR:                               gl.COMPRESSED_RGBA_ASTC_8x6_KHR,
		COMPRESSED_RGBA_ASTC_8x8_KHR:                               gl.COMPRESSED_RGBA_ASTC_8x8_KHR,
		COMPRESSED_RGBA_BPTC_UNORM:                                 gl.COMPRESSED_RGBA_BPTC_UNORM,
		COMPRESSED_RGBA_BPTC_UNORM_ARB:                             gl.COMPRESSED_RGBA_BPTC_UNORM_ARB,
		COMPRESSED_RGBA_S3TC_DXT1_EXT:                              gl.COMPRESSED_RGBA_S3TC_DXT1_EXT,
		COMPRESSED_RGBA_S3TC_DXT3_EXT:                              gl.COMPRESSED_RGBA_S3TC_DXT3_EXT,
		COMPRESSED_RGBA_S3TC_DXT5_EXT:                              gl.COMPRESSED_RGBA_S3TC_DXT5_EXT,
		COMPRESSED_RGB_BPTC_SIGNED_FLOAT:                           gl.COMPRESSED_RGB_BPTC_SIGNED_FLOAT,
		COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB:                       gl.COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB,
		COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT:                         gl.COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT,
		COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB:                     gl.COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB,
		COMPRESSED_RGB_S3TC_DXT1_EXT:                               gl.COMPRESSED_RGB_S3TC_DXT1_EXT,
		COMPRESSED_RG_RGTC2:                                        gl.COMPRESSED_RG_RGTC2,
		COMPRESSED_SIGNED_R11_EAC:                                  gl.COMPRESSED_SIGNED_R11_EAC,
		COMPRESSED_SIGNED_RED_RGTC1:                                gl.COMPRESSED_SIGNED_RED_RGTC1,
		COMPRESSED_SIGNED_RG11_EAC:                                 gl.COMPRESSED_SIGNED_RG11_EAC,
		COMPRESSED_SIGNED_RG_RGTC2:                                 gl.COMPRESSED_SIGNED_RG_RGTC2,
		COMPRESSED_SRGB:                                            gl.COMPRESSED_SRGB,
		COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR:                     gl.COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR:                      gl.COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR:                      gl.COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR:                      gl.COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR:                     gl.COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR:                     gl.COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR,
		COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR:                       gl.COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR,
		COMPRESSED_SRGB8_ALPHA8_ETC2_EAC:                           gl.COMPRESSED_SRGB8_ALPHA8_ETC2_EAC,
		COMPRESSED_SRGB8_ETC2:                                      gl.COMPRESSED_SRGB8_ETC2,
		COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:                  gl.COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2,
		COMPRESSED_SRGB_ALPHA:                                      gl.COMPRESSED_SRGB_ALPHA,
		COMPRESSED_SRGB_ALPHA_BPTC_UNORM:                           gl.COMPRESSED_SRGB_ALPHA_BPTC_UNORM,
		COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB:                       gl.COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB,
		COMPRESSED_TEXTURE_FORMATS:                                 gl.COMPRESSED_TEXTURE_FORMATS,
		COMPUTE_SHADER:                                             gl.COMPUTE_SHADER,
		COMPUTE_SHADER_BIT:                                         gl.COMPUTE_SHADER_BIT,
		COMPUTE_SHADER_INVOCATIONS:                                 gl.COMPUTE_SHADER_INVOCATIONS,
		COMPUTE_SHADER_INVOCATIONS_ARB:                             gl.COMPUTE_SHADER_INVOCATIONS_ARB,
		COMPUTE_SUBROUTINE:                                         gl.COMPUTE_SUBROUTINE,
		COMPUTE_SUBROUTINE_UNIFORM:                                 gl.COMPUTE_SUBROUTINE_UNIFORM,
		COMPUTE_TEXTURE:                                            gl.COMPUTE_TEXTURE,
		COMPUTE_WORK_GROUP_SIZE:                                    gl.COMPUTE_WORK_GROUP_SIZE,
		CONDITION_SATISFIED:                                        gl.CONDITION_SATISFIED,
		CONFORMANT_NV:                                              gl.CONFORMANT_NV,
		CONIC_CURVE_TO_NV:                                          gl.CONIC_CURVE_TO_NV,
		CONJOINT_NV:                                                gl.CONJOINT_NV,
		CONSERVATIVE_RASTERIZATION_INTEL:                           gl.CONSERVATIVE_RASTERIZATION_INTEL,
		CONSERVATIVE_RASTERIZATION_NV:                              gl.CONSERVATIVE_RASTERIZATION_NV,
		CONSERVATIVE_RASTER_DILATE_GRANULARITY_NV:                  gl.CONSERVATIVE_RASTER_DILATE_GRANULARITY_NV,
		CONSERVATIVE_RASTER_DILATE_NV:                              gl.CONSERVATIVE_RASTER_DILATE_NV,
		CONSERVATIVE_RASTER_DILATE_RANGE_NV:                        gl.CONSERVATIVE_RASTER_DILATE_RANGE_NV,
		CONSERVATIVE_RASTER_MODE_NV:                                gl.CONSERVATIVE_RASTER_MODE_NV,
		CONSERVATIVE_RASTER_MODE_POST_SNAP_NV:                      gl.CONSERVATIVE_RASTER_MODE_POST_SNAP_NV,
		CONSERVATIVE_RASTER_MODE_PRE_SNAP_NV:                       gl.CONSERVATIVE_RASTER_MODE_PRE_SNAP_NV,
		CONSERVATIVE_RASTER_MODE_PRE_SNAP_TRIANGLES_NV:             gl.CONSERVATIVE_RASTER_MODE_PRE_SNAP_TRIANGLES_NV,
		CONSTANT_ALPHA:                                             gl.CONSTANT_ALPHA,
		CONSTANT_COLOR:                                             gl.CONSTANT_COLOR,
		CONTEXT_COMPATIBILITY_PROFILE_BIT:                          gl.CONTEXT_COMPATIBILITY_PROFILE_BIT,
		CONTEXT_CORE_PROFILE_BIT:                                   gl.CONTEXT_CORE_PROFILE_BIT,
		CONTEXT_FLAGS:                                              gl.CONTEXT_FLAGS,
		CONTEXT_FLAG_DEBUG_BIT:                                     gl.CONTEXT_FLAG_DEBUG_BIT,
		CONTEXT_FLAG_DEBUG_BIT_KHR:                                 gl.CONTEXT_FLAG_DEBUG_BIT_KHR,
		CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT:                        gl.CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT,
		CONTEXT_FLAG_NO_ERROR_BIT:                                  gl.CONTEXT_FLAG_NO_ERROR_BIT,
		CONTEXT_FLAG_NO_ERROR_BIT_KHR:                              gl.CONTEXT_FLAG_NO_ERROR_BIT_KHR,
		CONTEXT_FLAG_ROBUST_ACCESS_BIT:                             gl.CONTEXT_FLAG_ROBUST_ACCESS_BIT,
		CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB:                         gl.CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB,
		CONTEXT_LOST:                                               gl.CONTEXT_LOST,
		CONTEXT_LOST_KHR:                                           gl.CONTEXT_LOST_KHR,
		CONTEXT_PROFILE_MASK:                                       gl.CONTEXT_PROFILE_MASK,
		CONTEXT_RELEASE_BEHAVIOR:                                   gl.CONTEXT_RELEASE_BEHAVIOR,
		CONTEXT_RELEASE_BEHAVIOR_FLUSH:                             gl.CONTEXT_RELEASE_BEHAVIOR_FLUSH,
		CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR:                         gl.CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR,
		CONTEXT_RELEASE_BEHAVIOR_KHR:                               gl.CONTEXT_RELEASE_BEHAVIOR_KHR,
		CONTEXT_ROBUST_ACCESS:                                      gl.CONTEXT_ROBUST_ACCESS,
		CONTEXT_ROBUST_ACCESS_KHR:                                  gl.CONTEXT_ROBUST_ACCESS_KHR,
		CONTRAST_NV:                                                gl.CONTRAST_NV,
		CONVEX_HULL_NV:                                             gl.CONVEX_HULL_NV,
		COPY:                                                       gl.COPY,
		COPY_INVERTED:                                              gl.COPY_INVERTED,
		COPY_READ_BUFFER:                                           gl.COPY_READ_BUFFER,
		COPY_READ_BUFFER_BINDING:                                   gl.COPY_READ_BUFFER_BINDING,
		COPY_WRITE_BUFFER:                                          gl.COPY_WRITE_BUFFER,
		COPY_WRITE_BUFFER_BINDING:                                  gl.COPY_WRITE_BUFFER_BINDING,
		COUNTER_RANGE_AMD:                                          gl.COUNTER_RANGE_AMD,
		COUNTER_TYPE_AMD:                                           gl.COUNTER_TYPE_AMD,
		COUNT_DOWN_NV:                                              gl.COUNT_DOWN_NV,
		COUNT_UP_NV:                                                gl.COUNT_UP_NV,
		COVERAGE_MODULATION_NV:                                     gl.COVERAGE_MODULATION_NV,
		COVERAGE_MODULATION_TABLE_NV:                               gl.COVERAGE_MODULATION_TABLE_NV,
		COVERAGE_MODULATION_TABLE_SIZE_NV:                          gl.COVERAGE_MODULATION_TABLE_SIZE_NV,
		CUBIC_CURVE_TO_NV:                                          gl.CUBIC_CURVE_TO_NV,
		CULL_FACE:                                                  gl.CULL_FACE,
		CULL_FACE_MODE:                                             gl.CULL_FACE_MODE,
		CURRENT_PROGRAM:                                            gl.CURRENT_PROGRAM,
		CURRENT_QUERY:                                              gl.CURRENT_QUERY,
		CURRENT_VERTEX_ATTRIB:                                      gl.CURRENT_VERTEX_ATTRIB,
		CW:                                                         gl.CW,
		DARKEN_KHR:                                                 gl.DARKEN_KHR,
		DARKEN_NV:                                                  gl.DARKEN_NV,
		DEBUG_CALLBACK_FUNCTION:                                    gl.DEBUG_CALLBACK_FUNCTION,
		DEBUG_CALLBACK_FUNCTION_ARB:                                gl.DEBUG_CALLBACK_FUNCTION_ARB,
		DEBUG_CALLBACK_FUNCTION_KHR:                                gl.DEBUG_CALLBACK_FUNCTION_KHR,
		DEBUG_CALLBACK_USER_PARAM:                                  gl.DEBUG_CALLBACK_USER_PARAM,
		DEBUG_CALLBACK_USER_PARAM_ARB:                              gl.DEBUG_CALLBACK_USER_PARAM_ARB,
		DEBUG_CALLBACK_USER_PARAM_KHR:                              gl.DEBUG_CALLBACK_USER_PARAM_KHR,
		DEBUG_GROUP_STACK_DEPTH:                                    gl.DEBUG_GROUP_STACK_DEPTH,
		DEBUG_GROUP_STACK_DEPTH_KHR:                                gl.DEBUG_GROUP_STACK_DEPTH_KHR,
		DEBUG_LOGGED_MESSAGES:                                      gl.DEBUG_LOGGED_MESSAGES,
		DEBUG_LOGGED_MESSAGES_ARB:                                  gl.DEBUG_LOGGED_MESSAGES_ARB,
		DEBUG_LOGGED_MESSAGES_KHR:                                  gl.DEBUG_LOGGED_MESSAGES_KHR,
		DEBUG_NEXT_LOGGED_MESSAGE_LENGTH:                           gl.DEBUG_NEXT_LOGGED_MESSAGE_LENGTH,
		DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB:                       gl.DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB,
		DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR:                       gl.DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR,
		DEBUG_OUTPUT:                                               gl.DEBUG_OUTPUT,
		DEBUG_OUTPUT_KHR:                                           gl.DEBUG_OUTPUT_KHR,
		DEBUG_OUTPUT_SYNCHRONOUS:                                   gl.DEBUG_OUTPUT_SYNCHRONOUS,
		DEBUG_OUTPUT_SYNCHRONOUS_ARB:                               gl.DEBUG_OUTPUT_SYNCHRONOUS_ARB,
		DEBUG_OUTPUT_SYNCHRONOUS_KHR:                               gl.DEBUG_OUTPUT_SYNCHRONOUS_KHR,
		DEBUG_SEVERITY_HIGH:                                        gl.DEBUG_SEVERITY_HIGH,
		DEBUG_SEVERITY_HIGH_ARB:                                    gl.DEBUG_SEVERITY_HIGH_ARB,
		DEBUG_SEVERITY_HIGH_KHR:                                    gl.DEBUG_SEVERITY_HIGH_KHR,
		DEBUG_SEVERITY_LOW:                                         gl.DEBUG_SEVERITY_LOW,
		DEBUG_SEVERITY_LOW_ARB:                                     gl.DEBUG_SEVERITY_LOW_ARB,
		DEBUG_SEVERITY_LOW_KHR:                                     gl.DEBUG_SEVERITY_LOW_KHR,
		DEBUG_SEVERITY_MEDIUM:                                      gl.DEBUG_SEVERITY_MEDIUM,
		DEBUG_SEVERITY_MEDIUM_ARB:                                  gl.DEBUG_SEVERITY_MEDIUM_ARB,
		DEBUG_SEVERITY_MEDIUM_KHR:                                  gl.DEBUG_SEVERITY_MEDIUM_KHR,
		DEBUG_SEVERITY_NOTIFICATION:                                gl.DEBUG_SEVERITY_NOTIFICATION,
		DEBUG_SEVERITY_NOTIFICATION_KHR:                            gl.DEBUG_SEVERITY_NOTIFICATION_KHR,
		DEBUG_SOURCE_API:                                           gl.DEBUG_SOURCE_API,
		DEBUG_SOURCE_API_ARB:                                       gl.DEBUG_SOURCE_API_ARB,
		DEBUG_SOURCE_API_KHR:                                       gl.DEBUG_SOURCE_API_KHR,
		DEBUG_SOURCE_APPLICATION:                                   gl.DEBUG_SOURCE_APPLICATION,
		DEBUG_SOURCE_APPLICATION_ARB:                               gl.DEBUG_SOURCE_APPLICATION_ARB,
		DEBUG_SOURCE_APPLICATION_KHR:                               gl.DEBUG_SOURCE_APPLICATION_KHR,
		DEBUG_SOURCE_OTHER:                                         gl.DEBUG_SOURCE_OTHER,
		DEBUG_SOURCE_OTHER_ARB:                                     gl.DEBUG_SOURCE_OTHER_ARB,
		DEBUG_SOURCE_OTHER_KHR:                                     gl.DEBUG_SOURCE_OTHER_KHR,
		DEBUG_SOURCE_SHADER_COMPILER:                               gl.DEBUG_SOURCE_SHADER_COMPILER,
		DEBUG_SOURCE_SHADER_COMPILER_ARB:                           gl.DEBUG_SOURCE_SHADER_COMPILER_ARB,
		DEBUG_SOURCE_SHADER_COMPILER_KHR:                           gl.DEBUG_SOURCE_SHADER_COMPILER_KHR,
		DEBUG_SOURCE_THIRD_PARTY:                                   gl.DEBUG_SOURCE_THIRD_PARTY,
		DEBUG_SOURCE_THIRD_PARTY_ARB:                               gl.DEBUG_SOURCE_THIRD_PARTY_ARB,
		DEBUG_SOURCE_THIRD_PARTY_KHR:                               gl.DEBUG_SOURCE_THIRD_PARTY_KHR,
		DEBUG_SOURCE_WINDOW_SYSTEM:                                 gl.DEBUG_SOURCE_WINDOW_SYSTEM,
		DEBUG_SOURCE_WINDOW_SYSTEM_ARB:                             gl.DEBUG_SOURCE_WINDOW_SYSTEM_ARB,
		DEBUG_SOURCE_WINDOW_SYSTEM_KHR:                             gl.DEBUG_SOURCE_WINDOW_SYSTEM_KHR,
		DEBUG_TYPE_DEPRECATED_BEHAVIOR:                             gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR,
		DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB:                         gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB,
		DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR:                         gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR,
		DEBUG_TYPE_ERROR:                                           gl.DEBUG_TYPE_ERROR,
		DEBUG_TYPE_ERROR_ARB:                                       gl.DEBUG_TYPE_ERROR_ARB,
		DEBUG_TYPE_ERROR_KHR:                                       gl.DEBUG_TYPE_ERROR_KHR,
		DEBUG_TYPE_MARKER:                                          gl.DEBUG_TYPE_MARKER,
		DEBUG_TYPE_MARKER_KHR:                                      gl.DEBUG_TYPE_MARKER_KHR,
		DEBUG_TYPE_OTHER:                                           gl.DEBUG_TYPE_OTHER,
		DEBUG_TYPE_OTHER_ARB:                                       gl.DEBUG_TYPE_OTHER_ARB,
		DEBUG_TYPE_OTHER_KHR:                                       gl.DEBUG_TYPE_OTHER_KHR,
		DEBUG_TYPE_PERFORMANCE:                                     gl.DEBUG_TYPE_PERFORMANCE,
		DEBUG_TYPE_PERFORMANCE_ARB:                                 gl.DEBUG_TYPE_PERFORMANCE_ARB,
		DEBUG_TYPE_PERFORMANCE_KHR:                                 gl.DEBUG_TYPE_PERFORMANCE_KHR,
		DEBUG_TYPE_POP_GROUP:                                       gl.DEBUG_TYPE_POP_GROUP,
		DEBUG_TYPE_POP_GROUP_KHR:                                   gl.DEBUG_TYPE_POP_GROUP_KHR,
		DEBUG_TYPE_PORTABILITY:                                     gl.DEBUG_TYPE_PORTABILITY,
		DEBUG_TYPE_PORTABILITY_ARB:                                 gl.DEBUG_TYPE_PORTABILITY_ARB,
		DEBUG_TYPE_PORTABILITY_KHR:                                 gl.DEBUG_TYPE_PORTABILITY_KHR,
		DEBUG_TYPE_PUSH_GROUP:                                      gl.DEBUG_TYPE_PUSH_GROUP,
		DEBUG_TYPE_PUSH_GROUP_KHR:                                  gl.DEBUG_TYPE_PUSH_GROUP_KHR,
		DEBUG_TYPE_UNDEFINED_BEHAVIOR:                              gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR,
		DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB:                          gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB,
		DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR:                          gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR,
		DECODE_EXT:                                                 gl.DECODE_EXT,
		DECR:                                                       gl.DECR,
		DECR_WRAP:                                                  gl.DECR_WRAP,
		DELETE_STATUS:                                              gl.DELETE_STATUS,
		DEPTH:                                                      gl.DEPTH,
		DEPTH24_STENCIL8:                                           gl.DEPTH24_STENCIL8,
		DEPTH32F_STENCIL8:                                          gl.DEPTH32F_STENCIL8,
		DEPTH_ATTACHMENT:                                           gl.DEPTH_ATTACHMENT,
		DEPTH_BUFFER_BIT:                                           gl.DEPTH_BUFFER_BIT,
		DEPTH_CLAMP:                                                gl.DEPTH_CLAMP,
		DEPTH_CLEAR_VALUE:                                          gl.DEPTH_CLEAR_VALUE,
		DEPTH_COMPONENT:                                            gl.DEPTH_COMPONENT,
		DEPTH_COMPONENT16:                                          gl.DEPTH_COMPONENT16,
		DEPTH_COMPONENT24:                                          gl.DEPTH_COMPONENT24,
		DEPTH_COMPONENT32:                                          gl.DEPTH_COMPONENT32,
		DEPTH_COMPONENT32F:                                         gl.DEPTH_COMPONENT32F,
		DEPTH_COMPONENTS:                                           gl.DEPTH_COMPONENTS,
		DEPTH_FUNC:                                                 gl.DEPTH_FUNC,
		DEPTH_RANGE:                                                gl.DEPTH_RANGE,
		DEPTH_RENDERABLE:                                           gl.DEPTH_RENDERABLE,
		DEPTH_SAMPLES_NV:                                           gl.DEPTH_SAMPLES_NV,
		DEPTH_STENCIL:                                              gl.DEPTH_STENCIL,
		DEPTH_STENCIL_ATTACHMENT:                                   gl.DEPTH_STENCIL_ATTACHMENT,
		DEPTH_STENCIL_TEXTURE_MODE:                                 gl.DEPTH_STENCIL_TEXTURE_MODE,
		DEPTH_TEST:                                                 gl.DEPTH_TEST,
		DEPTH_WRITEMASK:                                            gl.DEPTH_WRITEMASK,
		DIFFERENCE_KHR:                                             gl.DIFFERENCE_KHR,
		DIFFERENCE_NV:                                              gl.DIFFERENCE_NV,
		DISJOINT_NV:                                                gl.DISJOINT_NV,
		DISPATCH_INDIRECT_BUFFER:                                   gl.DISPATCH_INDIRECT_BUFFER,
		DISPATCH_INDIRECT_BUFFER_BINDING:                           gl.DISPATCH_INDIRECT_BUFFER_BINDING,
		DITHER:                                                     gl.DITHER,
		DONT_CARE:                                                  gl.DONT_CARE,
		DOUBLE:                                                     gl.DOUBLE,
		DOUBLEBUFFER:                                               gl.DOUBLEBUFFER,
		DOUBLE_MAT2:                                                gl.DOUBLE_MAT2,
		DOUBLE_MAT2x3:                                              gl.DOUBLE_MAT2x3,
		DOUBLE_MAT2x4:                                              gl.DOUBLE_MAT2x4,
		DOUBLE_MAT3:                                                gl.DOUBLE_MAT3,
		DOUBLE_MAT3x2:                                              gl.DOUBLE_MAT3x2,
		DOUBLE_MAT3x4:                                              gl.DOUBLE_MAT3x4,
		DOUBLE_MAT4:                                                gl.DOUBLE_MAT4,
		DOUBLE_MAT4x2:                                              gl.DOUBLE_MAT4x2,
		DOUBLE_MAT4x3:                                              gl.DOUBLE_MAT4x3,
		DOUBLE_VEC2:                                                gl.DOUBLE_VEC2,
		DOUBLE_VEC3:                                                gl.DOUBLE_VEC3,
		DOUBLE_VEC4:                                                gl.DOUBLE_VEC4,
		DRAW_ARRAYS_COMMAND_NV:                                     gl.DRAW_ARRAYS_COMMAND_NV,
		DRAW_ARRAYS_INSTANCED_COMMAND_NV:                           gl.DRAW_ARRAYS_INSTANCED_COMMAND_NV,
		DRAW_ARRAYS_STRIP_COMMAND_NV:                               gl.DRAW_ARRAYS_STRIP_COMMAND_NV,
		DRAW_BUFFER:                                                gl.DRAW_BUFFER,
		DRAW_BUFFER0:                                               gl.DRAW_BUFFER0,
		DRAW_BUFFER1:                                               gl.DRAW_BUFFER1,
		DRAW_BUFFER10:                                              gl.DRAW_BUFFER10,
		DRAW_BUFFER11:                                              gl.DRAW_BUFFER11,
		DRAW_BUFFER12:                                              gl.DRAW_BUFFER12,
		DRAW_BUFFER13:                                              gl.DRAW_BUFFER13,
		DRAW_BUFFER14:                                              gl.DRAW_BUFFER14,
		DRAW_BUFFER15:                                              gl.DRAW_BUFFER15,
		DRAW_BUFFER2:                                               gl.DRAW_BUFFER2,
		DRAW_BUFFER3:                                               gl.DRAW_BUFFER3,
		DRAW_BUFFER4:                                               gl.DRAW_BUFFER4,
		DRAW_BUFFER5:                                               gl.DRAW_BUFFER5,
		DRAW_BUFFER6:                                               gl.DRAW_BUFFER6,
		DRAW_BUFFER7:                                               gl.DRAW_BUFFER7,
		DRAW_BUFFER8:                                               gl.DRAW_BUFFER8,
		DRAW_BUFFER9:                                               gl.DRAW_BUFFER9,
		DRAW_ELEMENTS_COMMAND_NV:                                   gl.DRAW_ELEMENTS_COMMAND_NV,
		DRAW_ELEMENTS_INSTANCED_COMMAND_NV:                         gl.DRAW_ELEMENTS_INSTANCED_COMMAND_NV,
		DRAW_ELEMENTS_STRIP_COMMAND_NV:                             gl.DRAW_ELEMENTS_STRIP_COMMAND_NV,
		DRAW_FRAMEBUFFER:                                           gl.DRAW_FRAMEBUFFER,
		DRAW_FRAMEBUFFER_BINDING:                                   gl.DRAW_FRAMEBUFFER_BINDING,
		DRAW_INDIRECT_ADDRESS_NV:                                   gl.DRAW_INDIRECT_ADDRESS_NV,
		DRAW_INDIRECT_BUFFER:                                       gl.DRAW_INDIRECT_BUFFER,
		DRAW_INDIRECT_BUFFER_BINDING:                               gl.DRAW_INDIRECT_BUFFER_BINDING,
		DRAW_INDIRECT_LENGTH_NV:                                    gl.DRAW_INDIRECT_LENGTH_NV,
		DRAW_INDIRECT_UNIFIED_NV:                                   gl.DRAW_INDIRECT_UNIFIED_NV,
		DST_ALPHA:                                                  gl.DST_ALPHA,
		DST_ATOP_NV:                                                gl.DST_ATOP_NV,
		DST_COLOR:                                                  gl.DST_COLOR,
		DST_IN_NV:                                                  gl.DST_IN_NV,
		DST_NV:                                                     gl.DST_NV,
		DST_OUT_NV:                                                 gl.DST_OUT_NV,
		DST_OVER_NV:                                                gl.DST_OVER_NV,
		DUP_FIRST_CUBIC_CURVE_TO_NV:                                gl.DUP_FIRST_CUBIC_CURVE_TO_NV,
		DUP_LAST_CUBIC_CURVE_TO_NV:                                 gl.DUP_LAST_CUBIC_CURVE_TO_NV,
		DYNAMIC_COPY:                                               gl.DYNAMIC_COPY,
		DYNAMIC_DRAW:                                               gl.DYNAMIC_DRAW,
		DYNAMIC_READ:                                               gl.DYNAMIC_READ,
		DYNAMIC_STORAGE_BIT:                                        gl.DYNAMIC_STORAGE_BIT,
		EDGE_FLAG_ARRAY_ADDRESS_NV:                                 gl.EDGE_FLAG_ARRAY_ADDRESS_NV,
		EDGE_FLAG_ARRAY_LENGTH_NV:                                  gl.EDGE_FLAG_ARRAY_LENGTH_NV,
		EFFECTIVE_RASTER_SAMPLES_EXT:                               gl.EFFECTIVE_RASTER_SAMPLES_EXT,
		ELEMENT_ADDRESS_COMMAND_NV:                                 gl.ELEMENT_ADDRESS_COMMAND_NV,
		ELEMENT_ARRAY_ADDRESS_NV:                                   gl.ELEMENT_ARRAY_ADDRESS_NV,
		ELEMENT_ARRAY_BARRIER_BIT:                                  gl.ELEMENT_ARRAY_BARRIER_BIT,
		ELEMENT_ARRAY_BUFFER:                                       gl.ELEMENT_ARRAY_BUFFER,
		ELEMENT_ARRAY_BUFFER_BINDING:                               gl.ELEMENT_ARRAY_BUFFER_BINDING,
		ELEMENT_ARRAY_LENGTH_NV:                                    gl.ELEMENT_ARRAY_LENGTH_NV,
		ELEMENT_ARRAY_UNIFIED_NV:                                   gl.ELEMENT_ARRAY_UNIFIED_NV,
		EQUAL:                                                      gl.EQUAL,
		EQUIV:                                                      gl.EQUIV,
		EXCLUSION_KHR:                                              gl.EXCLUSION_KHR,
		EXCLUSION_NV:                                               gl.EXCLUSION_NV,
		EXCLUSIVE_EXT:                                              gl.EXCLUSIVE_EXT,
		EXTENSIONS:                                                 gl.EXTENSIONS,
		FACTOR_MAX_AMD:                                             gl.FACTOR_MAX_AMD,
		FACTOR_MIN_AMD:                                             gl.FACTOR_MIN_AMD,
		FALSE:                                                      gl.FALSE,
		FASTEST:                                                    gl.FASTEST,
		FILE_NAME_NV:                                               gl.FILE_NAME_NV,
		FILL:                                                       gl.FILL,
		FILL_RECTANGLE_NV:                                          gl.FILL_RECTANGLE_NV,
		FILTER:                                                     gl.FILTER,
		FIRST_TO_REST_NV:                                           gl.FIRST_TO_REST_NV,
		FIRST_VERTEX_CONVENTION:                                    gl.FIRST_VERTEX_CONVENTION,
		FIXED:                                                      gl.FIXED,
		FIXED_ONLY:                                                 gl.FIXED_ONLY,
		FLOAT:                                                      gl.FLOAT,
		FLOAT16_NV:                                                 gl.FLOAT16_NV,
		FLOAT16_VEC2_NV:                                            gl.FLOAT16_VEC2_NV,
		FLOAT16_VEC3_NV:                                            gl.FLOAT16_VEC3_NV,
		FLOAT16_VEC4_NV:                                            gl.FLOAT16_VEC4_NV,
		FLOAT_32_UNSIGNED_INT_24_8_REV:                             gl.FLOAT_32_UNSIGNED_INT_24_8_REV,
		FLOAT_MAT2:                                                 gl.FLOAT_MAT2,
		FLOAT_MAT2x3:                                               gl.FLOAT_MAT2x3,
		FLOAT_MAT2x4:                                               gl.FLOAT_MAT2x4,
		FLOAT_MAT3:                                                 gl.FLOAT_MAT3,
		FLOAT_MAT3x2:                                               gl.FLOAT_MAT3x2,
		FLOAT_MAT3x4:                                               gl.FLOAT_MAT3x4,
		FLOAT_MAT4:                                                 gl.FLOAT_MAT4,
		FLOAT_MAT4x2:                                               gl.FLOAT_MAT4x2,
		FLOAT_MAT4x3:                                               gl.FLOAT_MAT4x3,
		FLOAT_VEC2:                                                 gl.FLOAT_VEC2,
		FLOAT_VEC3:                                                 gl.FLOAT_VEC3,
		FLOAT_VEC4:                                                 gl.FLOAT_VEC4,
		FOG_COORD_ARRAY_ADDRESS_NV:                                 gl.FOG_COORD_ARRAY_ADDRESS_NV,
		FOG_COORD_ARRAY_LENGTH_NV:                                  gl.FOG_COORD_ARRAY_LENGTH_NV,
		FONT_ASCENDER_BIT_NV:                                       gl.FONT_ASCENDER_BIT_NV,
		FONT_DESCENDER_BIT_NV:                                      gl.FONT_DESCENDER_BIT_NV,
		FONT_GLYPHS_AVAILABLE_NV:                                   gl.FONT_GLYPHS_AVAILABLE_NV,
		FONT_HAS_KERNING_BIT_NV:                                    gl.FONT_HAS_KERNING_BIT_NV,
		FONT_HEIGHT_BIT_NV:                                         gl.FONT_HEIGHT_BIT_NV,
		FONT_MAX_ADVANCE_HEIGHT_BIT_NV:                             gl.FONT_MAX_ADVANCE_HEIGHT_BIT_NV,
		FONT_MAX_ADVANCE_WIDTH_BIT_NV:                              gl.FONT_MAX_ADVANCE_WIDTH_BIT_NV,
		FONT_NUM_GLYPH_INDICES_BIT_NV:                              gl.FONT_NUM_GLYPH_INDICES_BIT_NV,
		FONT_TARGET_UNAVAILABLE_NV:                                 gl.FONT_TARGET_UNAVAILABLE_NV,
		FONT_UNAVAILABLE_NV:                                        gl.FONT_UNAVAILABLE_NV,
		FONT_UNDERLINE_POSITION_BIT_NV:                             gl.FONT_UNDERLINE_POSITION_BIT_NV,
		FONT_UNDERLINE_THICKNESS_BIT_NV:                            gl.FONT_UNDERLINE_THICKNESS_BIT_NV,
		FONT_UNINTELLIGIBLE_NV:                                     gl.FONT_UNINTELLIGIBLE_NV,
		FONT_UNITS_PER_EM_BIT_NV:                                   gl.FONT_UNITS_PER_EM_BIT_NV,
		FONT_X_MAX_BOUNDS_BIT_NV:                                   gl.FONT_X_MAX_BOUNDS_BIT_NV,
		FONT_X_MIN_BOUNDS_BIT_NV:                                   gl.FONT_X_MIN_BOUNDS_BIT_NV,
		FONT_Y_MAX_BOUNDS_BIT_NV:                                   gl.FONT_Y_MAX_BOUNDS_BIT_NV,
		FONT_Y_MIN_BOUNDS_BIT_NV:                                   gl.FONT_Y_MIN_BOUNDS_BIT_NV,
		FRACTIONAL_EVEN:                                            gl.FRACTIONAL_EVEN,
		FRACTIONAL_ODD:                                             gl.FRACTIONAL_ODD,
		FRAGMENT_COVERAGE_COLOR_NV:                                 gl.FRAGMENT_COVERAGE_COLOR_NV,
		FRAGMENT_COVERAGE_TO_COLOR_NV:                              gl.FRAGMENT_COVERAGE_TO_COLOR_NV,
		FRAGMENT_INPUT_NV:                                          gl.FRAGMENT_INPUT_NV,
		FRAGMENT_INTERPOLATION_OFFSET_BITS:                         gl.FRAGMENT_INTERPOLATION_OFFSET_BITS,
		FRAGMENT_SHADER:                                            gl.FRAGMENT_SHADER,
		FRAGMENT_SHADER_BIT:                                        gl.FRAGMENT_SHADER_BIT,
		FRAGMENT_SHADER_BIT_EXT:                                    gl.FRAGMENT_SHADER_BIT_EXT,
		FRAGMENT_SHADER_DERIVATIVE_HINT:                            gl.FRAGMENT_SHADER_DERIVATIVE_HINT,
		FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT:                       gl.FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT,
		FRAGMENT_SHADER_INVOCATIONS:                                gl.FRAGMENT_SHADER_INVOCATIONS,
		FRAGMENT_SHADER_INVOCATIONS_ARB:                            gl.FRAGMENT_SHADER_INVOCATIONS_ARB,
		FRAGMENT_SUBROUTINE:                                        gl.FRAGMENT_SUBROUTINE,
		FRAGMENT_SUBROUTINE_UNIFORM:                                gl.FRAGMENT_SUBROUTINE_UNIFORM,
		FRAGMENT_TEXTURE:                                           gl.FRAGMENT_TEXTURE,
		FRAMEBUFFER:                                                gl.FRAMEBUFFER,
		FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE:                          gl.FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE,
		FRAMEBUFFER_ATTACHMENT_BLUE_SIZE:                           gl.FRAMEBUFFER_ATTACHMENT_BLUE_SIZE,
		FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING:                      gl.FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING,
		FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE:                      gl.FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE,
		FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE:                          gl.FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE,
		FRAMEBUFFER_ATTACHMENT_GREEN_SIZE:                          gl.FRAMEBUFFER_ATTACHMENT_GREEN_SIZE,
		FRAMEBUFFER_ATTACHMENT_LAYERED:                             gl.FRAMEBUFFER_ATTACHMENT_LAYERED,
		FRAMEBUFFER_ATTACHMENT_LAYERED_ARB:                         gl.FRAMEBUFFER_ATTACHMENT_LAYERED_ARB,
		FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:                         gl.FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,
		FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:                         gl.FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,
		FRAMEBUFFER_ATTACHMENT_RED_SIZE:                            gl.FRAMEBUFFER_ATTACHMENT_RED_SIZE,
		FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE:                        gl.FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR: gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:       gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:               gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:               gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,
		FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR:       gl.FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR,
		FRAMEBUFFER_BARRIER_BIT:                            gl.FRAMEBUFFER_BARRIER_BIT,
		FRAMEBUFFER_BINDING:                                gl.FRAMEBUFFER_BINDING,
		FRAMEBUFFER_BLEND:                                  gl.FRAMEBUFFER_BLEND,
		FRAMEBUFFER_COMPLETE:                               gl.FRAMEBUFFER_COMPLETE,
		FRAMEBUFFER_DEFAULT:                                gl.FRAMEBUFFER_DEFAULT,
		FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS:         gl.FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS,
		FRAMEBUFFER_DEFAULT_HEIGHT:                         gl.FRAMEBUFFER_DEFAULT_HEIGHT,
		FRAMEBUFFER_DEFAULT_LAYERS:                         gl.FRAMEBUFFER_DEFAULT_LAYERS,
		FRAMEBUFFER_DEFAULT_SAMPLES:                        gl.FRAMEBUFFER_DEFAULT_SAMPLES,
		FRAMEBUFFER_DEFAULT_WIDTH:                          gl.FRAMEBUFFER_DEFAULT_WIDTH,
		FRAMEBUFFER_INCOMPLETE_ATTACHMENT:                  gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT,
		FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:                 gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER,
		FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB:             gl.FRAMEBUFFER_INCOMPLETE_LAYER_COUNT_ARB,
		FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS:               gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS,
		FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB:           gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS_ARB,
		FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT:          gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,
		FRAMEBUFFER_INCOMPLETE_MULTISAMPLE:                 gl.FRAMEBUFFER_INCOMPLETE_MULTISAMPLE,
		FRAMEBUFFER_INCOMPLETE_READ_BUFFER:                 gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER,
		FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR:            gl.FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR,
		FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_ARB:      gl.FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_ARB,
		FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_NV:       gl.FRAMEBUFFER_PROGRAMMABLE_SAMPLE_LOCATIONS_NV,
		FRAMEBUFFER_RENDERABLE:                             gl.FRAMEBUFFER_RENDERABLE,
		FRAMEBUFFER_RENDERABLE_LAYERED:                     gl.FRAMEBUFFER_RENDERABLE_LAYERED,
		FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_ARB:         gl.FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_ARB,
		FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_NV:          gl.FRAMEBUFFER_SAMPLE_LOCATION_PIXEL_GRID_NV,
		FRAMEBUFFER_SRGB:                                   gl.FRAMEBUFFER_SRGB,
		FRAMEBUFFER_UNDEFINED:                              gl.FRAMEBUFFER_UNDEFINED,
		FRAMEBUFFER_UNSUPPORTED:                            gl.FRAMEBUFFER_UNSUPPORTED,
		FRONT:                                              gl.FRONT,
		FRONT_AND_BACK:                                     gl.FRONT_AND_BACK,
		FRONT_FACE:                                         gl.FRONT_FACE,
		FRONT_FACE_COMMAND_NV:                              gl.FRONT_FACE_COMMAND_NV,
		FRONT_LEFT:                                         gl.FRONT_LEFT,
		FRONT_RIGHT:                                        gl.FRONT_RIGHT,
		FULL_SUPPORT:                                       gl.FULL_SUPPORT,
		FUNC_ADD:                                           gl.FUNC_ADD,
		FUNC_REVERSE_SUBTRACT:                              gl.FUNC_REVERSE_SUBTRACT,
		FUNC_SUBTRACT:                                      gl.FUNC_SUBTRACT,
		GEOMETRY_INPUT_TYPE:                                gl.GEOMETRY_INPUT_TYPE,
		GEOMETRY_INPUT_TYPE_ARB:                            gl.GEOMETRY_INPUT_TYPE_ARB,
		GEOMETRY_OUTPUT_TYPE:                               gl.GEOMETRY_OUTPUT_TYPE,
		GEOMETRY_OUTPUT_TYPE_ARB:                           gl.GEOMETRY_OUTPUT_TYPE_ARB,
		GEOMETRY_SHADER:                                    gl.GEOMETRY_SHADER,
		GEOMETRY_SHADER_ARB:                                gl.GEOMETRY_SHADER_ARB,
		GEOMETRY_SHADER_BIT:                                gl.GEOMETRY_SHADER_BIT,
		GEOMETRY_SHADER_INVOCATIONS:                        gl.GEOMETRY_SHADER_INVOCATIONS,
		GEOMETRY_SHADER_PRIMITIVES_EMITTED:                 gl.GEOMETRY_SHADER_PRIMITIVES_EMITTED,
		GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB:             gl.GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB,
		GEOMETRY_SUBROUTINE:                                gl.GEOMETRY_SUBROUTINE,
		GEOMETRY_SUBROUTINE_UNIFORM:                        gl.GEOMETRY_SUBROUTINE_UNIFORM,
		GEOMETRY_TEXTURE:                                   gl.GEOMETRY_TEXTURE,
		GEOMETRY_VERTICES_OUT:                              gl.GEOMETRY_VERTICES_OUT,
		GEOMETRY_VERTICES_OUT_ARB:                          gl.GEOMETRY_VERTICES_OUT_ARB,
		GEQUAL:                                             gl.GEQUAL,
		GET_TEXTURE_IMAGE_FORMAT:                           gl.GET_TEXTURE_IMAGE_FORMAT,
		GET_TEXTURE_IMAGE_TYPE:                             gl.GET_TEXTURE_IMAGE_TYPE,
		GLYPH_HAS_KERNING_BIT_NV:                           gl.GLYPH_HAS_KERNING_BIT_NV,
		GLYPH_HEIGHT_BIT_NV:                                gl.GLYPH_HEIGHT_BIT_NV,
		GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV:            gl.GLYPH_HORIZONTAL_BEARING_ADVANCE_BIT_NV,
		GLYPH_HORIZONTAL_BEARING_X_BIT_NV:                  gl.GLYPH_HORIZONTAL_BEARING_X_BIT_NV,
		GLYPH_HORIZONTAL_BEARING_Y_BIT_NV:                  gl.GLYPH_HORIZONTAL_BEARING_Y_BIT_NV,
		GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV:              gl.GLYPH_VERTICAL_BEARING_ADVANCE_BIT_NV,
		GLYPH_VERTICAL_BEARING_X_BIT_NV:                    gl.GLYPH_VERTICAL_BEARING_X_BIT_NV,
		GLYPH_VERTICAL_BEARING_Y_BIT_NV:                    gl.GLYPH_VERTICAL_BEARING_Y_BIT_NV,
		GLYPH_WIDTH_BIT_NV:                                 gl.GLYPH_WIDTH_BIT_NV,
		GPU_ADDRESS_NV:                                     gl.GPU_ADDRESS_NV,
		GREATER:                                            gl.GREATER,
		GREEN:                                              gl.GREEN,
		GREEN_INTEGER:                                      gl.GREEN_INTEGER,
		GREEN_NV:                                           gl.GREEN_NV,
		GUILTY_CONTEXT_RESET:                               gl.GUILTY_CONTEXT_RESET,
		GUILTY_CONTEXT_RESET_ARB:                           gl.GUILTY_CONTEXT_RESET_ARB,
		GUILTY_CONTEXT_RESET_KHR:                           gl.GUILTY_CONTEXT_RESET_KHR,
		HALF_FLOAT:                                         gl.HALF_FLOAT,
		HARDLIGHT_KHR:                                      gl.HARDLIGHT_KHR,
		HARDLIGHT_NV:                                       gl.HARDLIGHT_NV,
		HARDMIX_NV:                                         gl.HARDMIX_NV,
		HIGH_FLOAT:                                         gl.HIGH_FLOAT,
		HIGH_INT:                                           gl.HIGH_INT,
		HORIZONTAL_LINE_TO_NV:                              gl.HORIZONTAL_LINE_TO_NV,
		HSL_COLOR_KHR:                                      gl.HSL_COLOR_KHR,
		HSL_COLOR_NV:                                       gl.HSL_COLOR_NV,
		HSL_HUE_KHR:                                        gl.HSL_HUE_KHR,
		HSL_HUE_NV:                                         gl.HSL_HUE_NV,
		HSL_LUMINOSITY_KHR:                                 gl.HSL_LUMINOSITY_KHR,
		HSL_LUMINOSITY_NV:                                  gl.HSL_LUMINOSITY_NV,
		HSL_SATURATION_KHR:                                 gl.HSL_SATURATION_KHR,
		HSL_SATURATION_NV:                                  gl.HSL_SATURATION_NV,
		IMAGE_1D:                                           gl.IMAGE_1D,
		IMAGE_1D_ARRAY:                                     gl.IMAGE_1D_ARRAY,
		IMAGE_2D:                                           gl.IMAGE_2D,
		IMAGE_2D_ARRAY:                                     gl.IMAGE_2D_ARRAY,
		IMAGE_2D_MULTISAMPLE:                               gl.IMAGE_2D_MULTISAMPLE,
		IMAGE_2D_MULTISAMPLE_ARRAY:                         gl.IMAGE_2D_MULTISAMPLE_ARRAY,
		IMAGE_2D_RECT:                                      gl.IMAGE_2D_RECT,
		IMAGE_3D:                                           gl.IMAGE_3D,
		IMAGE_BINDING_ACCESS:                               gl.IMAGE_BINDING_ACCESS,
		IMAGE_BINDING_FORMAT:                               gl.IMAGE_BINDING_FORMAT,
		IMAGE_BINDING_LAYER:                                gl.IMAGE_BINDING_LAYER,
		IMAGE_BINDING_LAYERED:                              gl.IMAGE_BINDING_LAYERED,
		IMAGE_BINDING_LEVEL:                                gl.IMAGE_BINDING_LEVEL,
		IMAGE_BINDING_NAME:                                 gl.IMAGE_BINDING_NAME,
		IMAGE_BUFFER:                                       gl.IMAGE_BUFFER,
		IMAGE_CLASS_10_10_10_2:                             gl.IMAGE_CLASS_10_10_10_2,
		IMAGE_CLASS_11_11_10:                               gl.IMAGE_CLASS_11_11_10,
		IMAGE_CLASS_1_X_16:                                 gl.IMAGE_CLASS_1_X_16,
		IMAGE_CLASS_1_X_32:                                 gl.IMAGE_CLASS_1_X_32,
		IMAGE_CLASS_1_X_8:                                  gl.IMAGE_CLASS_1_X_8,
		IMAGE_CLASS_2_X_16:                                 gl.IMAGE_CLASS_2_X_16,
		IMAGE_CLASS_2_X_32:                                 gl.IMAGE_CLASS_2_X_32,
		IMAGE_CLASS_2_X_8:                                  gl.IMAGE_CLASS_2_X_8,
		IMAGE_CLASS_4_X_16:                                 gl.IMAGE_CLASS_4_X_16,
		IMAGE_CLASS_4_X_32:                                 gl.IMAGE_CLASS_4_X_32,
		IMAGE_CLASS_4_X_8:                                  gl.IMAGE_CLASS_4_X_8,
		IMAGE_COMPATIBILITY_CLASS:                          gl.IMAGE_COMPATIBILITY_CLASS,
		IMAGE_CUBE:                                         gl.IMAGE_CUBE,
		IMAGE_CUBE_MAP_ARRAY:                               gl.IMAGE_CUBE_MAP_ARRAY,
		IMAGE_FORMAT_COMPATIBILITY_BY_CLASS:                gl.IMAGE_FORMAT_COMPATIBILITY_BY_CLASS,
		IMAGE_FORMAT_COMPATIBILITY_BY_SIZE:                 gl.IMAGE_FORMAT_COMPATIBILITY_BY_SIZE,
		IMAGE_FORMAT_COMPATIBILITY_TYPE:                    gl.IMAGE_FORMAT_COMPATIBILITY_TYPE,
		IMAGE_PIXEL_FORMAT:                                 gl.IMAGE_PIXEL_FORMAT,
		IMAGE_PIXEL_TYPE:                                   gl.IMAGE_PIXEL_TYPE,
		IMAGE_TEXEL_SIZE:                                   gl.IMAGE_TEXEL_SIZE,
		IMPLEMENTATION_COLOR_READ_FORMAT:                   gl.IMPLEMENTATION_COLOR_READ_FORMAT,
		IMPLEMENTATION_COLOR_READ_TYPE:                     gl.IMPLEMENTATION_COLOR_READ_TYPE,
		INCLUSIVE_EXT:                                      gl.INCLUSIVE_EXT,
		INCR:                                               gl.INCR,
		INCR_WRAP:                                          gl.INCR_WRAP,
		INDEX_ARRAY_ADDRESS_NV:                             gl.INDEX_ARRAY_ADDRESS_NV,
		INDEX_ARRAY_LENGTH_NV:                              gl.INDEX_ARRAY_LENGTH_NV,
		INFO_LOG_LENGTH:                                    gl.INFO_LOG_LENGTH,
		INNOCENT_CONTEXT_RESET:                             gl.INNOCENT_CONTEXT_RESET,
		INNOCENT_CONTEXT_RESET_ARB:                         gl.INNOCENT_CONTEXT_RESET_ARB,
		INNOCENT_CONTEXT_RESET_KHR:                         gl.INNOCENT_CONTEXT_RESET_KHR,
		INT:                                                gl.INT,
		INT16_NV:                                           gl.INT16_NV,
		INT16_VEC2_NV:                                      gl.INT16_VEC2_NV,
		INT16_VEC3_NV:                                      gl.INT16_VEC3_NV,
		INT16_VEC4_NV:                                      gl.INT16_VEC4_NV,
		INT64_ARB:                                          gl.INT64_ARB,
		INT64_NV:                                           gl.INT64_NV,
		INT64_VEC2_ARB:                                     gl.INT64_VEC2_ARB,
		INT64_VEC2_NV:                                      gl.INT64_VEC2_NV,
		INT64_VEC3_ARB:                                     gl.INT64_VEC3_ARB,
		INT64_VEC3_NV:                                      gl.INT64_VEC3_NV,
		INT64_VEC4_ARB:                                     gl.INT64_VEC4_ARB,
		INT64_VEC4_NV:                                      gl.INT64_VEC4_NV,
		INT8_NV:                                            gl.INT8_NV,
		INT8_VEC2_NV:                                       gl.INT8_VEC2_NV,
		INT8_VEC3_NV:                                       gl.INT8_VEC3_NV,
		INT8_VEC4_NV:                                       gl.INT8_VEC4_NV,
		INTERLEAVED_ATTRIBS:                                gl.INTERLEAVED_ATTRIBS,
		INTERNALFORMAT_ALPHA_SIZE:                          gl.INTERNALFORMAT_ALPHA_SIZE,
		INTERNALFORMAT_ALPHA_TYPE:                          gl.INTERNALFORMAT_ALPHA_TYPE,
		INTERNALFORMAT_BLUE_SIZE:                           gl.INTERNALFORMAT_BLUE_SIZE,
		INTERNALFORMAT_BLUE_TYPE:                           gl.INTERNALFORMAT_BLUE_TYPE,
		INTERNALFORMAT_DEPTH_SIZE:                          gl.INTERNALFORMAT_DEPTH_SIZE,
		INTERNALFORMAT_DEPTH_TYPE:                          gl.INTERNALFORMAT_DEPTH_TYPE,
		INTERNALFORMAT_GREEN_SIZE:                          gl.INTERNALFORMAT_GREEN_SIZE,
		INTERNALFORMAT_GREEN_TYPE:                          gl.INTERNALFORMAT_GREEN_TYPE,
		INTERNALFORMAT_PREFERRED:                           gl.INTERNALFORMAT_PREFERRED,
		INTERNALFORMAT_RED_SIZE:                            gl.INTERNALFORMAT_RED_SIZE,
		INTERNALFORMAT_RED_TYPE:                            gl.INTERNALFORMAT_RED_TYPE,
		INTERNALFORMAT_SHARED_SIZE:                         gl.INTERNALFORMAT_SHARED_SIZE,
		INTERNALFORMAT_STENCIL_SIZE:                        gl.INTERNALFORMAT_STENCIL_SIZE,
		INTERNALFORMAT_STENCIL_TYPE:                        gl.INTERNALFORMAT_STENCIL_TYPE,
		INTERNALFORMAT_SUPPORTED:                           gl.INTERNALFORMAT_SUPPORTED,
		INT_2_10_10_10_REV:                                 gl.INT_2_10_10_10_REV,
		INT_IMAGE_1D:                                       gl.INT_IMAGE_1D,
		INT_IMAGE_1D_ARRAY:                                 gl.INT_IMAGE_1D_ARRAY,
		INT_IMAGE_2D:                                       gl.INT_IMAGE_2D,
		INT_IMAGE_2D_ARRAY:                                 gl.INT_IMAGE_2D_ARRAY,
		INT_IMAGE_2D_MULTISAMPLE:                           gl.INT_IMAGE_2D_MULTISAMPLE,
		INT_IMAGE_2D_MULTISAMPLE_ARRAY:                     gl.INT_IMAGE_2D_MULTISAMPLE_ARRAY,
		INT_IMAGE_2D_RECT:                                  gl.INT_IMAGE_2D_RECT,
		INT_IMAGE_3D:                                       gl.INT_IMAGE_3D,
		INT_IMAGE_BUFFER:                                   gl.INT_IMAGE_BUFFER,
		INT_IMAGE_CUBE:                                     gl.INT_IMAGE_CUBE,
		INT_IMAGE_CUBE_MAP_ARRAY:                           gl.INT_IMAGE_CUBE_MAP_ARRAY,
		INT_SAMPLER_1D:                                     gl.INT_SAMPLER_1D,
		INT_SAMPLER_1D_ARRAY:                               gl.INT_SAMPLER_1D_ARRAY,
		INT_SAMPLER_2D:                                     gl.INT_SAMPLER_2D,
		INT_SAMPLER_2D_ARRAY:                               gl.INT_SAMPLER_2D_ARRAY,
		INT_SAMPLER_2D_MULTISAMPLE:                         gl.INT_SAMPLER_2D_MULTISAMPLE,
		INT_SAMPLER_2D_MULTISAMPLE_ARRAY:                   gl.INT_SAMPLER_2D_MULTISAMPLE_ARRAY,
		INT_SAMPLER_2D_RECT:                                gl.INT_SAMPLER_2D_RECT,
		INT_SAMPLER_3D:                                     gl.INT_SAMPLER_3D,
		INT_SAMPLER_BUFFER:                                 gl.INT_SAMPLER_BUFFER,
		INT_SAMPLER_CUBE:                                   gl.INT_SAMPLER_CUBE,
		INT_SAMPLER_CUBE_MAP_ARRAY:                         gl.INT_SAMPLER_CUBE_MAP_ARRAY,
		INT_SAMPLER_CUBE_MAP_ARRAY_ARB:                     gl.INT_SAMPLER_CUBE_MAP_ARRAY_ARB,
		INT_VEC2:                                           gl.INT_VEC2,
		INT_VEC3:                                           gl.INT_VEC3,
		INT_VEC4:                                           gl.INT_VEC4,
		INVALID_ENUM:                                       gl.INVALID_ENUM,
		INVALID_FRAMEBUFFER_OPERATION:                      gl.INVALID_FRAMEBUFFER_OPERATION,
		INVALID_INDEX:                                      gl.INVALID_INDEX,
		INVALID_OPERATION:                                  gl.INVALID_OPERATION,
		INVALID_VALUE:                                      gl.INVALID_VALUE,
		INVERT:                                             gl.INVERT,
		INVERT_OVG_NV:                                      gl.INVERT_OVG_NV,
		INVERT_RGB_NV:                                      gl.INVERT_RGB_NV,
		ISOLINES:                                           gl.ISOLINES,
		IS_PER_PATCH:                                       gl.IS_PER_PATCH,
		IS_ROW_MAJOR:                                       gl.IS_ROW_MAJOR,
		ITALIC_BIT_NV:                                      gl.ITALIC_BIT_NV,
		KEEP:                                               gl.KEEP,
		LARGE_CCW_ARC_TO_NV:                                gl.LARGE_CCW_ARC_TO_NV,
		LARGE_CW_ARC_TO_NV:                                 gl.LARGE_CW_ARC_TO_NV,
		LAST_VERTEX_CONVENTION:                             gl.LAST_VERTEX_CONVENTION,
		LAYER_PROVOKING_VERTEX:                             gl.LAYER_PROVOKING_VERTEX,
		LEFT:                                               gl.LEFT,
		LEQUAL:                                             gl.LEQUAL,
		LESS:                                               gl.LESS,
		LIGHTEN_KHR:                                        gl.LIGHTEN_KHR,
		LIGHTEN_NV:                                         gl.LIGHTEN_NV,
		LINE:                                               gl.LINE,
		LINEAR:                                             gl.LINEAR,
		LINEARBURN_NV:                                      gl.LINEARBURN_NV,
		LINEARDODGE_NV:                                     gl.LINEARDODGE_NV,
		LINEARLIGHT_NV:                                     gl.LINEARLIGHT_NV,
		LINEAR_MIPMAP_LINEAR:                               gl.LINEAR_MIPMAP_LINEAR,
		LINEAR_MIPMAP_NEAREST:                              gl.LINEAR_MIPMAP_NEAREST,
		LINES:                                              gl.LINES,
		LINES_ADJACENCY:                                    gl.LINES_ADJACENCY,
		LINES_ADJACENCY_ARB:                                gl.LINES_ADJACENCY_ARB,
		LINE_LOOP:                                          gl.LINE_LOOP,
		LINE_SMOOTH:                                        gl.LINE_SMOOTH,
		LINE_SMOOTH_HINT:                                   gl.LINE_SMOOTH_HINT,
		LINE_STRIP:                                         gl.LINE_STRIP,
		LINE_STRIP_ADJACENCY:                               gl.LINE_STRIP_ADJACENCY,
		LINE_STRIP_ADJACENCY_ARB:                           gl.LINE_STRIP_ADJACENCY_ARB,
		LINE_TO_NV:                                         gl.LINE_TO_NV,
		LINE_WIDTH:                                         gl.LINE_WIDTH,
		LINE_WIDTH_COMMAND_NV:                              gl.LINE_WIDTH_COMMAND_NV,
		LINE_WIDTH_GRANULARITY:                             gl.LINE_WIDTH_GRANULARITY,
		LINE_WIDTH_RANGE:                                   gl.LINE_WIDTH_RANGE,
		LINK_STATUS:                                        gl.LINK_STATUS,
		LOCATION:                                           gl.LOCATION,
		LOCATION_COMPONENT:                                 gl.LOCATION_COMPONENT,
		LOCATION_INDEX:                                     gl.LOCATION_INDEX,
		LOGIC_OP_MODE:                                      gl.LOGIC_OP_MODE,
		LOSE_CONTEXT_ON_RESET:                              gl.LOSE_CONTEXT_ON_RESET,
		LOSE_CONTEXT_ON_RESET_ARB:                          gl.LOSE_CONTEXT_ON_RESET_ARB,
		LOSE_CONTEXT_ON_RESET_KHR:                          gl.LOSE_CONTEXT_ON_RESET_KHR,
		LOWER_LEFT:                                         gl.LOWER_LEFT,
		LOW_FLOAT:                                          gl.LOW_FLOAT,
		LOW_INT:                                            gl.LOW_INT,
		MAJOR_VERSION:                                      gl.MAJOR_VERSION,
		MANUAL_GENERATE_MIPMAP:                             gl.MANUAL_GENERATE_MIPMAP,
		MAP_COHERENT_BIT:                                   gl.MAP_COHERENT_BIT,
		MAP_FLUSH_EXPLICIT_BIT:                             gl.MAP_FLUSH_EXPLICIT_BIT,
		MAP_INVALIDATE_BUFFER_BIT:                          gl.MAP_INVALIDATE_BUFFER_BIT,
		MAP_INVALIDATE_RANGE_BIT:                           gl.MAP_INVALIDATE_RANGE_BIT,
		MAP_PERSISTENT_BIT:                                 gl.MAP_PERSISTENT_BIT,
		MAP_READ_BIT:                                       gl.MAP_READ_BIT,
		MAP_UNSYNCHRONIZED_BIT:                             gl.MAP_UNSYNCHRONIZED_BIT,
		MAP_WRITE_BIT:                                      gl.MAP_WRITE_BIT,
		MATRIX_STRIDE:                                      gl.MATRIX_STRIDE,
		MAX:                                                gl.MAX,
		MAX_3D_TEXTURE_SIZE:                                gl.MAX_3D_TEXTURE_SIZE,
		MAX_ARRAY_TEXTURE_LAYERS:                           gl.MAX_ARRAY_TEXTURE_LAYERS,
		MAX_ATOMIC_COUNTER_BUFFER_BINDINGS:                 gl.MAX_ATOMIC_COUNTER_BUFFER_BINDINGS,
		MAX_ATOMIC_COUNTER_BUFFER_SIZE:                     gl.MAX_ATOMIC_COUNTER_BUFFER_SIZE,
		MAX_CLIP_DISTANCES:                                 gl.MAX_CLIP_DISTANCES,
		MAX_COLOR_ATTACHMENTS:                              gl.MAX_COLOR_ATTACHMENTS,
		MAX_COLOR_TEXTURE_SAMPLES:                          gl.MAX_COLOR_TEXTURE_SAMPLES,
		MAX_COMBINED_ATOMIC_COUNTERS:                       gl.MAX_COMBINED_ATOMIC_COUNTERS,
		MAX_COMBINED_ATOMIC_COUNTER_BUFFERS:                gl.MAX_COMBINED_ATOMIC_COUNTER_BUFFERS,
		MAX_COMBINED_CLIP_AND_CULL_DISTANCES:               gl.MAX_COMBINED_CLIP_AND_CULL_DISTANCES,
		MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS:            gl.MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS,
		MAX_COMBINED_DIMENSIONS:                            gl.MAX_COMBINED_DIMENSIONS,
		MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS:           gl.MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS,
		MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS:           gl.MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS,
		MAX_COMBINED_IMAGE_UNIFORMS:                        gl.MAX_COMBINED_IMAGE_UNIFORMS,
		MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS:   gl.MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS,
		MAX_COMBINED_SHADER_OUTPUT_RESOURCES:            gl.MAX_COMBINED_SHADER_OUTPUT_RESOURCES,
		MAX_COMBINED_SHADER_STORAGE_BLOCKS:              gl.MAX_COMBINED_SHADER_STORAGE_BLOCKS,
		MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS:    gl.MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS,
		MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS: gl.MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS,
		MAX_COMBINED_TEXTURE_IMAGE_UNITS:                gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS,
		MAX_COMBINED_UNIFORM_BLOCKS:                     gl.MAX_COMBINED_UNIFORM_BLOCKS,
		MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS:          gl.MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS,
		MAX_COMPUTE_ATOMIC_COUNTERS:                     gl.MAX_COMPUTE_ATOMIC_COUNTERS,
		MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS:              gl.MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS,
		MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB:         gl.MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB,
		MAX_COMPUTE_FIXED_GROUP_SIZE_ARB:                gl.MAX_COMPUTE_FIXED_GROUP_SIZE_ARB,
		MAX_COMPUTE_IMAGE_UNIFORMS:                      gl.MAX_COMPUTE_IMAGE_UNIFORMS,
		MAX_COMPUTE_SHADER_STORAGE_BLOCKS:               gl.MAX_COMPUTE_SHADER_STORAGE_BLOCKS,
		MAX_COMPUTE_SHARED_MEMORY_SIZE:                  gl.MAX_COMPUTE_SHARED_MEMORY_SIZE,
		MAX_COMPUTE_TEXTURE_IMAGE_UNITS:                 gl.MAX_COMPUTE_TEXTURE_IMAGE_UNITS,
		MAX_COMPUTE_UNIFORM_BLOCKS:                      gl.MAX_COMPUTE_UNIFORM_BLOCKS,
		MAX_COMPUTE_UNIFORM_COMPONENTS:                  gl.MAX_COMPUTE_UNIFORM_COMPONENTS,
		MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB:      gl.MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB,
		MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB:             gl.MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB,
		MAX_COMPUTE_WORK_GROUP_COUNT:                    gl.MAX_COMPUTE_WORK_GROUP_COUNT,
		MAX_COMPUTE_WORK_GROUP_INVOCATIONS:              gl.MAX_COMPUTE_WORK_GROUP_INVOCATIONS,
		MAX_COMPUTE_WORK_GROUP_SIZE:                     gl.MAX_COMPUTE_WORK_GROUP_SIZE,
		MAX_CUBE_MAP_TEXTURE_SIZE:                       gl.MAX_CUBE_MAP_TEXTURE_SIZE,
		MAX_CULL_DISTANCES:                              gl.MAX_CULL_DISTANCES,
		MAX_DEBUG_GROUP_STACK_DEPTH:                     gl.MAX_DEBUG_GROUP_STACK_DEPTH,
		MAX_DEBUG_GROUP_STACK_DEPTH_KHR:                 gl.MAX_DEBUG_GROUP_STACK_DEPTH_KHR,
		MAX_DEBUG_LOGGED_MESSAGES:                       gl.MAX_DEBUG_LOGGED_MESSAGES,
		MAX_DEBUG_LOGGED_MESSAGES_ARB:                   gl.MAX_DEBUG_LOGGED_MESSAGES_ARB,
		MAX_DEBUG_LOGGED_MESSAGES_KHR:                   gl.MAX_DEBUG_LOGGED_MESSAGES_KHR,
		MAX_DEBUG_MESSAGE_LENGTH:                        gl.MAX_DEBUG_MESSAGE_LENGTH,
		MAX_DEBUG_MESSAGE_LENGTH_ARB:                    gl.MAX_DEBUG_MESSAGE_LENGTH_ARB,
		MAX_DEBUG_MESSAGE_LENGTH_KHR:                    gl.MAX_DEBUG_MESSAGE_LENGTH_KHR,
		MAX_DEPTH:                                       gl.MAX_DEPTH,
		MAX_DEPTH_TEXTURE_SAMPLES:                       gl.MAX_DEPTH_TEXTURE_SAMPLES,
		MAX_DRAW_BUFFERS:                                gl.MAX_DRAW_BUFFERS,
		MAX_DUAL_SOURCE_DRAW_BUFFERS:                    gl.MAX_DUAL_SOURCE_DRAW_BUFFERS,
		MAX_ELEMENTS_INDICES:                            gl.MAX_ELEMENTS_INDICES,
		MAX_ELEMENTS_VERTICES:                           gl.MAX_ELEMENTS_VERTICES,
		MAX_ELEMENT_INDEX:                               gl.MAX_ELEMENT_INDEX,
		MAX_FRAGMENT_ATOMIC_COUNTERS:                    gl.MAX_FRAGMENT_ATOMIC_COUNTERS,
		MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS:             gl.MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS,
		MAX_FRAGMENT_IMAGE_UNIFORMS:                     gl.MAX_FRAGMENT_IMAGE_UNIFORMS,
		MAX_FRAGMENT_INPUT_COMPONENTS:                   gl.MAX_FRAGMENT_INPUT_COMPONENTS,
		MAX_FRAGMENT_INTERPOLATION_OFFSET:               gl.MAX_FRAGMENT_INTERPOLATION_OFFSET,
		MAX_FRAGMENT_SHADER_STORAGE_BLOCKS:              gl.MAX_FRAGMENT_SHADER_STORAGE_BLOCKS,
		MAX_FRAGMENT_UNIFORM_BLOCKS:                     gl.MAX_FRAGMENT_UNIFORM_BLOCKS,
		MAX_FRAGMENT_UNIFORM_COMPONENTS:                 gl.MAX_FRAGMENT_UNIFORM_COMPONENTS,
		MAX_FRAGMENT_UNIFORM_VECTORS:                    gl.MAX_FRAGMENT_UNIFORM_VECTORS,
		MAX_FRAMEBUFFER_HEIGHT:                          gl.MAX_FRAMEBUFFER_HEIGHT,
		MAX_FRAMEBUFFER_LAYERS:                          gl.MAX_FRAMEBUFFER_LAYERS,
		MAX_FRAMEBUFFER_SAMPLES:                         gl.MAX_FRAMEBUFFER_SAMPLES,
		MAX_FRAMEBUFFER_WIDTH:                           gl.MAX_FRAMEBUFFER_WIDTH,
		MAX_GEOMETRY_ATOMIC_COUNTERS:                    gl.MAX_GEOMETRY_ATOMIC_COUNTERS,
		MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS:             gl.MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS,
		MAX_GEOMETRY_IMAGE_UNIFORMS:                     gl.MAX_GEOMETRY_IMAGE_UNIFORMS,
		MAX_GEOMETRY_INPUT_COMPONENTS:                   gl.MAX_GEOMETRY_INPUT_COMPONENTS,
		MAX_GEOMETRY_OUTPUT_COMPONENTS:                  gl.MAX_GEOMETRY_OUTPUT_COMPONENTS,
		MAX_GEOMETRY_OUTPUT_VERTICES:                    gl.MAX_GEOMETRY_OUTPUT_VERTICES,
		MAX_GEOMETRY_OUTPUT_VERTICES_ARB:                gl.MAX_GEOMETRY_OUTPUT_VERTICES_ARB,
		MAX_GEOMETRY_SHADER_INVOCATIONS:                 gl.MAX_GEOMETRY_SHADER_INVOCATIONS,
		MAX_GEOMETRY_SHADER_STORAGE_BLOCKS:              gl.MAX_GEOMETRY_SHADER_STORAGE_BLOCKS,
		MAX_GEOMETRY_TEXTURE_IMAGE_UNITS:                gl.MAX_GEOMETRY_TEXTURE_IMAGE_UNITS,
		MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB:            gl.MAX_GEOMETRY_TEXTURE_IMAGE_UNITS_ARB,
		MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS:            gl.MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS,
		MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB:        gl.MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS_ARB,
		MAX_GEOMETRY_UNIFORM_BLOCKS:                     gl.MAX_GEOMETRY_UNIFORM_BLOCKS,
		MAX_GEOMETRY_UNIFORM_COMPONENTS:                 gl.MAX_GEOMETRY_UNIFORM_COMPONENTS,
		MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB:             gl.MAX_GEOMETRY_UNIFORM_COMPONENTS_ARB,
		MAX_GEOMETRY_VARYING_COMPONENTS_ARB:             gl.MAX_GEOMETRY_VARYING_COMPONENTS_ARB,
		MAX_HEIGHT:                                      gl.MAX_HEIGHT,
		MAX_IMAGE_SAMPLES:                               gl.MAX_IMAGE_SAMPLES,
		MAX_IMAGE_UNITS:                                 gl.MAX_IMAGE_UNITS,
		MAX_INTEGER_SAMPLES:                             gl.MAX_INTEGER_SAMPLES,
		MAX_LABEL_LENGTH:                                gl.MAX_LABEL_LENGTH,
		MAX_LABEL_LENGTH_KHR:                            gl.MAX_LABEL_LENGTH_KHR,
		MAX_LAYERS:                                      gl.MAX_LAYERS,
		MAX_MULTISAMPLE_COVERAGE_MODES_NV:               gl.MAX_MULTISAMPLE_COVERAGE_MODES_NV,
		MAX_NAME_LENGTH:                                 gl.MAX_NAME_LENGTH,
		MAX_NUM_ACTIVE_VARIABLES:                        gl.MAX_NUM_ACTIVE_VARIABLES,
		MAX_NUM_COMPATIBLE_SUBROUTINES:                  gl.MAX_NUM_COMPATIBLE_SUBROUTINES,
		MAX_PATCH_VERTICES:                              gl.MAX_PATCH_VERTICES,
		MAX_PROGRAM_TEXEL_OFFSET:                        gl.MAX_PROGRAM_TEXEL_OFFSET,
		MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB:       gl.MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB,
		MAX_PROGRAM_TEXTURE_GATHER_OFFSET:               gl.MAX_PROGRAM_TEXTURE_GATHER_OFFSET,
		MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB:           gl.MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB,
		MAX_RASTER_SAMPLES_EXT:                          gl.MAX_RASTER_SAMPLES_EXT,
		MAX_RECTANGLE_TEXTURE_SIZE:                      gl.MAX_RECTANGLE_TEXTURE_SIZE,
		MAX_RENDERBUFFER_SIZE:                           gl.MAX_RENDERBUFFER_SIZE,
		MAX_SAMPLES:                                     gl.MAX_SAMPLES,
		MAX_SAMPLE_MASK_WORDS:                           gl.MAX_SAMPLE_MASK_WORDS,
		MAX_SERVER_WAIT_TIMEOUT:                         gl.MAX_SERVER_WAIT_TIMEOUT,
		MAX_SHADER_BUFFER_ADDRESS_NV:                    gl.MAX_SHADER_BUFFER_ADDRESS_NV,
		MAX_SHADER_COMPILER_THREADS_ARB:                 gl.MAX_SHADER_COMPILER_THREADS_ARB,
		MAX_SHADER_COMPILER_THREADS_KHR:                 gl.MAX_SHADER_COMPILER_THREADS_KHR,
		MAX_SHADER_STORAGE_BLOCK_SIZE:                   gl.MAX_SHADER_STORAGE_BLOCK_SIZE,
		MAX_SHADER_STORAGE_BUFFER_BINDINGS:              gl.MAX_SHADER_STORAGE_BUFFER_BINDINGS,
		MAX_SPARSE_3D_TEXTURE_SIZE_ARB:                  gl.MAX_SPARSE_3D_TEXTURE_SIZE_ARB,
		MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB:             gl.MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB,
		MAX_SPARSE_TEXTURE_SIZE_ARB:                     gl.MAX_SPARSE_TEXTURE_SIZE_ARB,
		MAX_SUBPIXEL_PRECISION_BIAS_BITS_NV:             gl.MAX_SUBPIXEL_PRECISION_BIAS_BITS_NV,
		MAX_SUBROUTINES:                                 gl.MAX_SUBROUTINES,
		MAX_SUBROUTINE_UNIFORM_LOCATIONS:                gl.MAX_SUBROUTINE_UNIFORM_LOCATIONS,
		MAX_TESS_CONTROL_ATOMIC_COUNTERS:                gl.MAX_TESS_CONTROL_ATOMIC_COUNTERS,
		MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS:         gl.MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS,
		MAX_TESS_CONTROL_IMAGE_UNIFORMS:                 gl.MAX_TESS_CONTROL_IMAGE_UNIFORMS,
		MAX_TESS_CONTROL_INPUT_COMPONENTS:               gl.MAX_TESS_CONTROL_INPUT_COMPONENTS,
		MAX_TESS_CONTROL_OUTPUT_COMPONENTS:              gl.MAX_TESS_CONTROL_OUTPUT_COMPONENTS,
		MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS:          gl.MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS,
		MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS:            gl.MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS,
		MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS:        gl.MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS,
		MAX_TESS_CONTROL_UNIFORM_BLOCKS:                 gl.MAX_TESS_CONTROL_UNIFORM_BLOCKS,
		MAX_TESS_CONTROL_UNIFORM_COMPONENTS:             gl.MAX_TESS_CONTROL_UNIFORM_COMPONENTS,
		MAX_TESS_EVALUATION_ATOMIC_COUNTERS:             gl.MAX_TESS_EVALUATION_ATOMIC_COUNTERS,
		MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS:      gl.MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS,
		MAX_TESS_EVALUATION_IMAGE_UNIFORMS:              gl.MAX_TESS_EVALUATION_IMAGE_UNIFORMS,
		MAX_TESS_EVALUATION_INPUT_COMPONENTS:            gl.MAX_TESS_EVALUATION_INPUT_COMPONENTS,
		MAX_TESS_EVALUATION_OUTPUT_COMPONENTS:           gl.MAX_TESS_EVALUATION_OUTPUT_COMPONENTS,
		MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS:       gl.MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS,
		MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS:         gl.MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS,
		MAX_TESS_EVALUATION_UNIFORM_BLOCKS:              gl.MAX_TESS_EVALUATION_UNIFORM_BLOCKS,
		MAX_TESS_EVALUATION_UNIFORM_COMPONENTS:          gl.MAX_TESS_EVALUATION_UNIFORM_COMPONENTS,
		MAX_TESS_GEN_LEVEL:                              gl.MAX_TESS_GEN_LEVEL,
		MAX_TESS_PATCH_COMPONENTS:                       gl.MAX_TESS_PATCH_COMPONENTS,
		MAX_TEXTURE_BUFFER_SIZE:                         gl.MAX_TEXTURE_BUFFER_SIZE,
		MAX_TEXTURE_BUFFER_SIZE_ARB:                     gl.MAX_TEXTURE_BUFFER_SIZE_ARB,
		MAX_TEXTURE_IMAGE_UNITS:                         gl.MAX_TEXTURE_IMAGE_UNITS,
		MAX_TEXTURE_LOD_BIAS:                            gl.MAX_TEXTURE_LOD_BIAS,
		MAX_TEXTURE_MAX_ANISOTROPY:                      gl.MAX_TEXTURE_MAX_ANISOTROPY,
		MAX_TEXTURE_SIZE:                                gl.MAX_TEXTURE_SIZE,
		MAX_TRANSFORM_FEEDBACK_BUFFERS:                  gl.MAX_TRANSFORM_FEEDBACK_BUFFERS,
		MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS:   gl.MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS,
		MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS:         gl.MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS,
		MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS:      gl.MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS,
		MAX_UNIFORM_BLOCK_SIZE:                          gl.MAX_UNIFORM_BLOCK_SIZE,
		MAX_UNIFORM_BUFFER_BINDINGS:                     gl.MAX_UNIFORM_BUFFER_BINDINGS,
		MAX_UNIFORM_LOCATIONS:                           gl.MAX_UNIFORM_LOCATIONS,
		MAX_VARYING_COMPONENTS:                          gl.MAX_VARYING_COMPONENTS,
		MAX_VARYING_FLOATS:                              gl.MAX_VARYING_FLOATS,
		MAX_VARYING_VECTORS:                             gl.MAX_VARYING_VECTORS,
		MAX_VERTEX_ATOMIC_COUNTERS:                      gl.MAX_VERTEX_ATOMIC_COUNTERS,
		MAX_VERTEX_ATOMIC_COUNTER_BUFFERS:               gl.MAX_VERTEX_ATOMIC_COUNTER_BUFFERS,
		MAX_VERTEX_ATTRIBS:                              gl.MAX_VERTEX_ATTRIBS,
		MAX_VERTEX_ATTRIB_BINDINGS:                      gl.MAX_VERTEX_ATTRIB_BINDINGS,
		MAX_VERTEX_ATTRIB_RELATIVE_OFFSET:               gl.MAX_VERTEX_ATTRIB_RELATIVE_OFFSET,
		MAX_VERTEX_ATTRIB_STRIDE:                        gl.MAX_VERTEX_ATTRIB_STRIDE,
		MAX_VERTEX_IMAGE_UNIFORMS:                       gl.MAX_VERTEX_IMAGE_UNIFORMS,
		MAX_VERTEX_OUTPUT_COMPONENTS:                    gl.MAX_VERTEX_OUTPUT_COMPONENTS,
		MAX_VERTEX_SHADER_STORAGE_BLOCKS:                gl.MAX_VERTEX_SHADER_STORAGE_BLOCKS,
		MAX_VERTEX_STREAMS:                              gl.MAX_VERTEX_STREAMS,
		MAX_VERTEX_TEXTURE_IMAGE_UNITS:                  gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS,
		MAX_VERTEX_UNIFORM_BLOCKS:                       gl.MAX_VERTEX_UNIFORM_BLOCKS,
		MAX_VERTEX_UNIFORM_COMPONENTS:                   gl.MAX_VERTEX_UNIFORM_COMPONENTS,
		MAX_VERTEX_UNIFORM_VECTORS:                      gl.MAX_VERTEX_UNIFORM_VECTORS,
		MAX_VERTEX_VARYING_COMPONENTS_ARB:               gl.MAX_VERTEX_VARYING_COMPONENTS_ARB,
		MAX_VIEWPORTS:                                   gl.MAX_VIEWPORTS,
		MAX_VIEWPORT_DIMS:                               gl.MAX_VIEWPORT_DIMS,
		MAX_VIEWS_OVR:                                   gl.MAX_VIEWS_OVR,
		MAX_WIDTH:                                       gl.MAX_WIDTH,
		MAX_WINDOW_RECTANGLES_EXT:                       gl.MAX_WINDOW_RECTANGLES_EXT,
		MEDIUM_FLOAT:                                    gl.MEDIUM_FLOAT,
		MEDIUM_INT:                                      gl.MEDIUM_INT,
		MIN:                                             gl.MIN,
		MINOR_VERSION:                                   gl.MINOR_VERSION,
		MINUS_CLAMPED_NV:                                gl.MINUS_CLAMPED_NV,
		MINUS_NV:                                        gl.MINUS_NV,
		MIN_FRAGMENT_INTERPOLATION_OFFSET:               gl.MIN_FRAGMENT_INTERPOLATION_OFFSET,
		MIN_MAP_BUFFER_ALIGNMENT:                        gl.MIN_MAP_BUFFER_ALIGNMENT,
		MIN_PROGRAM_TEXEL_OFFSET:                        gl.MIN_PROGRAM_TEXEL_OFFSET,
		MIN_PROGRAM_TEXTURE_GATHER_OFFSET:               gl.MIN_PROGRAM_TEXTURE_GATHER_OFFSET,
		MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB:           gl.MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB,
		MIN_SAMPLE_SHADING_VALUE:                        gl.MIN_SAMPLE_SHADING_VALUE,
		MIN_SAMPLE_SHADING_VALUE_ARB:                    gl.MIN_SAMPLE_SHADING_VALUE_ARB,
		MIPMAP:                                          gl.MIPMAP,
		MIRRORED_REPEAT:                                 gl.MIRRORED_REPEAT,
		MIRRORED_REPEAT_ARB:                             gl.MIRRORED_REPEAT_ARB,
		MIRROR_CLAMP_TO_EDGE:                            gl.MIRROR_CLAMP_TO_EDGE,
		MITER_REVERT_NV:                                 gl.MITER_REVERT_NV,
		MITER_TRUNCATE_NV:                               gl.MITER_TRUNCATE_NV,
		MIXED_DEPTH_SAMPLES_SUPPORTED_NV:                gl.MIXED_DEPTH_SAMPLES_SUPPORTED_NV,
		MIXED_STENCIL_SAMPLES_SUPPORTED_NV:              gl.MIXED_STENCIL_SAMPLES_SUPPORTED_NV,
		MOVE_TO_CONTINUES_NV:                            gl.MOVE_TO_CONTINUES_NV,
		MOVE_TO_NV:                                      gl.MOVE_TO_NV,
		MOVE_TO_RESETS_NV:                               gl.MOVE_TO_RESETS_NV,
		MULTIPLY_KHR:                                    gl.MULTIPLY_KHR,
		MULTIPLY_NV:                                     gl.MULTIPLY_NV,
		MULTISAMPLE:                                     gl.MULTISAMPLE,
		MULTISAMPLES_NV:                                 gl.MULTISAMPLES_NV,
		MULTISAMPLE_COVERAGE_MODES_NV:                   gl.MULTISAMPLE_COVERAGE_MODES_NV,
		MULTISAMPLE_LINE_WIDTH_GRANULARITY_ARB:          gl.MULTISAMPLE_LINE_WIDTH_GRANULARITY_ARB,
		MULTISAMPLE_LINE_WIDTH_RANGE_ARB:                gl.MULTISAMPLE_LINE_WIDTH_RANGE_ARB,
		MULTISAMPLE_RASTERIZATION_ALLOWED_EXT:           gl.MULTISAMPLE_RASTERIZATION_ALLOWED_EXT,
		NAMED_STRING_LENGTH_ARB:                         gl.NAMED_STRING_LENGTH_ARB,
		NAMED_STRING_TYPE_ARB:                           gl.NAMED_STRING_TYPE_ARB,
		NAME_LENGTH:                                     gl.NAME_LENGTH,
		NAND:                                            gl.NAND,
		NEAREST:                                         gl.NEAREST,
		NEAREST_MIPMAP_LINEAR:                           gl.NEAREST_MIPMAP_LINEAR,
		NEAREST_MIPMAP_NEAREST:                          gl.NEAREST_MIPMAP_NEAREST,
		NEGATIVE_ONE_TO_ONE:                             gl.NEGATIVE_ONE_TO_ONE,
		NEVER:                                           gl.NEVER,
		NICEST:                                          gl.NICEST,
		NONE:                                            gl.NONE,
		NOOP:                                            gl.NOOP,
		NOP_COMMAND_NV:                                  gl.NOP_COMMAND_NV,
		NOR:                                             gl.NOR,
		NORMAL_ARRAY_ADDRESS_NV:                         gl.NORMAL_ARRAY_ADDRESS_NV,
		NORMAL_ARRAY_LENGTH_NV:                          gl.NORMAL_ARRAY_LENGTH_NV,
		NOTEQUAL:                                        gl.NOTEQUAL,
		NO_ERROR:                                        gl.NO_ERROR,
		NO_RESET_NOTIFICATION:                           gl.NO_RESET_NOTIFICATION,
		NO_RESET_NOTIFICATION_ARB:                       gl.NO_RESET_NOTIFICATION_ARB,
		NO_RESET_NOTIFICATION_KHR:                       gl.NO_RESET_NOTIFICATION_KHR,
		NUM_ACTIVE_VARIABLES:                            gl.NUM_ACTIVE_VARIABLES,
		NUM_COMPATIBLE_SUBROUTINES:                      gl.NUM_COMPATIBLE_SUBROUTINES,
		NUM_COMPRESSED_TEXTURE_FORMATS:                  gl.NUM_COMPRESSED_TEXTURE_FORMATS,
		NUM_EXTENSIONS:                                  gl.NUM_EXTENSIONS,
		NUM_PROGRAM_BINARY_FORMATS:                      gl.NUM_PROGRAM_BINARY_FORMATS,
		NUM_SAMPLE_COUNTS:                               gl.NUM_SAMPLE_COUNTS,
		NUM_SHADER_BINARY_FORMATS:                       gl.NUM_SHADER_BINARY_FORMATS,
		NUM_SHADING_LANGUAGE_VERSIONS:                   gl.NUM_SHADING_LANGUAGE_VERSIONS,
		NUM_SPARSE_LEVELS_ARB:                           gl.NUM_SPARSE_LEVELS_ARB,
		NUM_SPIR_V_EXTENSIONS:                           gl.NUM_SPIR_V_EXTENSIONS,
		NUM_VIRTUAL_PAGE_SIZES_ARB:                      gl.NUM_VIRTUAL_PAGE_SIZES_ARB,
		NUM_WINDOW_RECTANGLES_EXT:                       gl.NUM_WINDOW_RECTANGLES_EXT,
		OBJECT_TYPE:                                     gl.OBJECT_TYPE,
		OFFSET:                                          gl.OFFSET,
		ONE:                                             gl.ONE,
		ONE_MINUS_CONSTANT_ALPHA:                        gl.ONE_MINUS_CONSTANT_ALPHA,
		ONE_MINUS_CONSTANT_COLOR:                        gl.ONE_MINUS_CONSTANT_COLOR,
		ONE_MINUS_DST_ALPHA:                             gl.ONE_MINUS_DST_ALPHA,
		ONE_MINUS_DST_COLOR:                             gl.ONE_MINUS_DST_COLOR,
		ONE_MINUS_SRC1_ALPHA:                            gl.ONE_MINUS_SRC1_ALPHA,
		ONE_MINUS_SRC1_COLOR:                            gl.ONE_MINUS_SRC1_COLOR,
		ONE_MINUS_SRC_ALPHA:                             gl.ONE_MINUS_SRC_ALPHA,
		ONE_MINUS_SRC_COLOR:                             gl.ONE_MINUS_SRC_COLOR,
		OR:                                              gl.OR,
		OR_INVERTED:                                     gl.OR_INVERTED,
		OR_REVERSE:                                      gl.OR_REVERSE,
		OUT_OF_MEMORY:                                   gl.OUT_OF_MEMORY,
		OVERLAY_KHR:                                     gl.OVERLAY_KHR,
		OVERLAY_NV:                                      gl.OVERLAY_NV,
		PACK_ALIGNMENT:                                  gl.PACK_ALIGNMENT,
		PACK_COMPRESSED_BLOCK_DEPTH:                     gl.PACK_COMPRESSED_BLOCK_DEPTH,
		PACK_COMPRESSED_BLOCK_HEIGHT:                    gl.PACK_COMPRESSED_BLOCK_HEIGHT,
		PACK_COMPRESSED_BLOCK_SIZE:                      gl.PACK_COMPRESSED_BLOCK_SIZE,
		PACK_COMPRESSED_BLOCK_WIDTH:                     gl.PACK_COMPRESSED_BLOCK_WIDTH,
		PACK_IMAGE_HEIGHT:                               gl.PACK_IMAGE_HEIGHT,
		PACK_LSB_FIRST:                                  gl.PACK_LSB_FIRST,
		PACK_ROW_LENGTH:                                 gl.PACK_ROW_LENGTH,
		PACK_SKIP_IMAGES:                                gl.PACK_SKIP_IMAGES,
		PACK_SKIP_PIXELS:                                gl.PACK_SKIP_PIXELS,
		PACK_SKIP_ROWS:                                  gl.PACK_SKIP_ROWS,
		PACK_SWAP_BYTES:                                 gl.PACK_SWAP_BYTES,
		PARAMETER_BUFFER:                                gl.PARAMETER_BUFFER,
		PARAMETER_BUFFER_ARB:                            gl.PARAMETER_BUFFER_ARB,
		PARAMETER_BUFFER_BINDING:                        gl.PARAMETER_BUFFER_BINDING,
		PARAMETER_BUFFER_BINDING_ARB:                    gl.PARAMETER_BUFFER_BINDING_ARB,
		PATCHES:                                         gl.PATCHES,
		PATCH_DEFAULT_INNER_LEVEL:                       gl.PATCH_DEFAULT_INNER_LEVEL,
		PATCH_DEFAULT_OUTER_LEVEL:                       gl.PATCH_DEFAULT_OUTER_LEVEL,
		PATCH_VERTICES:                                  gl.PATCH_VERTICES,
		PATH_CLIENT_LENGTH_NV:                           gl.PATH_CLIENT_LENGTH_NV,
		PATH_COMMAND_COUNT_NV:                           gl.PATH_COMMAND_COUNT_NV,
		PATH_COMPUTED_LENGTH_NV:                         gl.PATH_COMPUTED_LENGTH_NV,
		PATH_COORD_COUNT_NV:                             gl.PATH_COORD_COUNT_NV,
		PATH_COVER_DEPTH_FUNC_NV:                        gl.PATH_COVER_DEPTH_FUNC_NV,
		PATH_DASH_ARRAY_COUNT_NV:                        gl.PATH_DASH_ARRAY_COUNT_NV,
		PATH_DASH_CAPS_NV:                               gl.PATH_DASH_CAPS_NV,
		PATH_DASH_OFFSET_NV:                             gl.PATH_DASH_OFFSET_NV,
		PATH_DASH_OFFSET_RESET_NV:                       gl.PATH_DASH_OFFSET_RESET_NV,
		PATH_END_CAPS_NV:                                gl.PATH_END_CAPS_NV,
		PATH_ERROR_POSITION_NV:                          gl.PATH_ERROR_POSITION_NV,
		PATH_FILL_BOUNDING_BOX_NV:                       gl.PATH_FILL_BOUNDING_BOX_NV,
		PATH_FILL_COVER_MODE_NV:                         gl.PATH_FILL_COVER_MODE_NV,
		PATH_FILL_MASK_NV:                               gl.PATH_FILL_MASK_NV,
		PATH_FILL_MODE_NV:                               gl.PATH_FILL_MODE_NV,
		PATH_FORMAT_PS_NV:                               gl.PATH_FORMAT_PS_NV,
		PATH_FORMAT_SVG_NV:                              gl.PATH_FORMAT_SVG_NV,
		PATH_GEN_COEFF_NV:                               gl.PATH_GEN_COEFF_NV,
		PATH_GEN_COMPONENTS_NV:                          gl.PATH_GEN_COMPONENTS_NV,
		PATH_GEN_MODE_NV:                                gl.PATH_GEN_MODE_NV,
		PATH_INITIAL_DASH_CAP_NV:                        gl.PATH_INITIAL_DASH_CAP_NV,
		PATH_INITIAL_END_CAP_NV:                         gl.PATH_INITIAL_END_CAP_NV,
		PATH_JOIN_STYLE_NV:                              gl.PATH_JOIN_STYLE_NV,
		PATH_MAX_MODELVIEW_STACK_DEPTH_NV:               gl.PATH_MAX_MODELVIEW_STACK_DEPTH_NV,
		PATH_MAX_PROJECTION_STACK_DEPTH_NV:              gl.PATH_MAX_PROJECTION_STACK_DEPTH_NV,
		PATH_MITER_LIMIT_NV:                             gl.PATH_MITER_LIMIT_NV,
		PATH_MODELVIEW_MATRIX_NV:                        gl.PATH_MODELVIEW_MATRIX_NV,
		PATH_MODELVIEW_NV:                               gl.PATH_MODELVIEW_NV,
		PATH_MODELVIEW_STACK_DEPTH_NV:                   gl.PATH_MODELVIEW_STACK_DEPTH_NV,
		PATH_OBJECT_BOUNDING_BOX_NV:                     gl.PATH_OBJECT_BOUNDING_BOX_NV,
		PATH_PROJECTION_MATRIX_NV:                       gl.PATH_PROJECTION_MATRIX_NV,
		PATH_PROJECTION_NV:                              gl.PATH_PROJECTION_NV,
		PATH_PROJECTION_STACK_DEPTH_NV:                  gl.PATH_PROJECTION_STACK_DEPTH_NV,
		PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV:             gl.PATH_STENCIL_DEPTH_OFFSET_FACTOR_NV,
		PATH_STENCIL_DEPTH_OFFSET_UNITS_NV:              gl.PATH_STENCIL_DEPTH_OFFSET_UNITS_NV,
		PATH_STENCIL_FUNC_NV:                            gl.PATH_STENCIL_FUNC_NV,
		PATH_STENCIL_REF_NV:                             gl.PATH_STENCIL_REF_NV,
		PATH_STENCIL_VALUE_MASK_NV:                      gl.PATH_STENCIL_VALUE_MASK_NV,
		PATH_STROKE_BOUNDING_BOX_NV:                     gl.PATH_STROKE_BOUNDING_BOX_NV,
		PATH_STROKE_COVER_MODE_NV:                       gl.PATH_STROKE_COVER_MODE_NV,
		PATH_STROKE_MASK_NV:                             gl.PATH_STROKE_MASK_NV,
		PATH_STROKE_WIDTH_NV:                            gl.PATH_STROKE_WIDTH_NV,
		PATH_TERMINAL_DASH_CAP_NV:                       gl.PATH_TERMINAL_DASH_CAP_NV,
		PATH_TERMINAL_END_CAP_NV:                        gl.PATH_TERMINAL_END_CAP_NV,
		PATH_TRANSPOSE_MODELVIEW_MATRIX_NV:              gl.PATH_TRANSPOSE_MODELVIEW_MATRIX_NV,
		PATH_TRANSPOSE_PROJECTION_MATRIX_NV:             gl.PATH_TRANSPOSE_PROJECTION_MATRIX_NV,
		PERCENTAGE_AMD:                                  gl.PERCENTAGE_AMD,
		PERFMON_RESULT_AMD:                              gl.PERFMON_RESULT_AMD,
		PERFMON_RESULT_AVAILABLE_AMD:                    gl.PERFMON_RESULT_AVAILABLE_AMD,
		PERFMON_RESULT_SIZE_AMD:                         gl.PERFMON_RESULT_SIZE_AMD,
		PERFQUERY_COUNTER_DATA_BOOL32_INTEL:             gl.PERFQUERY_COUNTER_DATA_BOOL32_INTEL,
		PERFQUERY_COUNTER_DATA_DOUBLE_INTEL:             gl.PERFQUERY_COUNTER_DATA_DOUBLE_INTEL,
		PERFQUERY_COUNTER_DATA_FLOAT_INTEL:              gl.PERFQUERY_COUNTER_DATA_FLOAT_INTEL,
		PERFQUERY_COUNTER_DATA_UINT32_INTEL:             gl.PERFQUERY_COUNTER_DATA_UINT32_INTEL,
		PERFQUERY_COUNTER_DATA_UINT64_INTEL:             gl.PERFQUERY_COUNTER_DATA_UINT64_INTEL,
		PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL:         gl.PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL,
		PERFQUERY_COUNTER_DURATION_NORM_INTEL:           gl.PERFQUERY_COUNTER_DURATION_NORM_INTEL,
		PERFQUERY_COUNTER_DURATION_RAW_INTEL:            gl.PERFQUERY_COUNTER_DURATION_RAW_INTEL,
		PERFQUERY_COUNTER_EVENT_INTEL:                   gl.PERFQUERY_COUNTER_EVENT_INTEL,
		PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL:         gl.PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL,
		PERFQUERY_COUNTER_RAW_INTEL:                     gl.PERFQUERY_COUNTER_RAW_INTEL,
		PERFQUERY_COUNTER_THROUGHPUT_INTEL:              gl.PERFQUERY_COUNTER_THROUGHPUT_INTEL,
		PERFQUERY_COUNTER_TIMESTAMP_INTEL:               gl.PERFQUERY_COUNTER_TIMESTAMP_INTEL,
		PERFQUERY_DONOT_FLUSH_INTEL:                     gl.PERFQUERY_DONOT_FLUSH_INTEL,
		PERFQUERY_FLUSH_INTEL:                           gl.PERFQUERY_FLUSH_INTEL,
		PERFQUERY_GLOBAL_CONTEXT_INTEL:                  gl.PERFQUERY_GLOBAL_CONTEXT_INTEL,
		PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL:           gl.PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL,
		PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL:           gl.PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL,
		PERFQUERY_SINGLE_CONTEXT_INTEL:                  gl.PERFQUERY_SINGLE_CONTEXT_INTEL,
		PERFQUERY_WAIT_INTEL:                            gl.PERFQUERY_WAIT_INTEL,
		PINLIGHT_NV:                                     gl.PINLIGHT_NV,
		PIXEL_BUFFER_BARRIER_BIT:                        gl.PIXEL_BUFFER_BARRIER_BIT,
		PIXEL_PACK_BUFFER:                               gl.PIXEL_PACK_BUFFER,
		PIXEL_PACK_BUFFER_ARB:                           gl.PIXEL_PACK_BUFFER_ARB,
		PIXEL_PACK_BUFFER_BINDING:                       gl.PIXEL_PACK_BUFFER_BINDING,
		PIXEL_PACK_BUFFER_BINDING_ARB:                   gl.PIXEL_PACK_BUFFER_BINDING_ARB,
		PIXEL_UNPACK_BUFFER:                             gl.PIXEL_UNPACK_BUFFER,
		PIXEL_UNPACK_BUFFER_ARB:                         gl.PIXEL_UNPACK_BUFFER_ARB,
		PIXEL_UNPACK_BUFFER_BINDING:                     gl.PIXEL_UNPACK_BUFFER_BINDING,
		PIXEL_UNPACK_BUFFER_BINDING_ARB:                 gl.PIXEL_UNPACK_BUFFER_BINDING_ARB,
		PLUS_CLAMPED_ALPHA_NV:                           gl.PLUS_CLAMPED_ALPHA_NV,
		PLUS_CLAMPED_NV:                                 gl.PLUS_CLAMPED_NV,
		PLUS_DARKER_NV:                                  gl.PLUS_DARKER_NV,
		PLUS_NV:                                         gl.PLUS_NV,
		POINT:                                           gl.POINT,
		POINTS:                                          gl.POINTS,
		POINT_FADE_THRESHOLD_SIZE:                       gl.POINT_FADE_THRESHOLD_SIZE,
		POINT_SIZE:                                      gl.POINT_SIZE,
		POINT_SIZE_GRANULARITY:                          gl.POINT_SIZE_GRANULARITY,
		POINT_SIZE_RANGE:                                gl.POINT_SIZE_RANGE,
		POINT_SPRITE_COORD_ORIGIN:                       gl.POINT_SPRITE_COORD_ORIGIN,
		POLYGON_MODE:                                    gl.POLYGON_MODE,
		POLYGON_OFFSET_CLAMP:                            gl.POLYGON_OFFSET_CLAMP,
		POLYGON_OFFSET_CLAMP_EXT:                        gl.POLYGON_OFFSET_CLAMP_EXT,
		POLYGON_OFFSET_COMMAND_NV:                       gl.POLYGON_OFFSET_COMMAND_NV,
		POLYGON_OFFSET_FACTOR:                           gl.POLYGON_OFFSET_FACTOR,
		POLYGON_OFFSET_FILL:                             gl.POLYGON_OFFSET_FILL,
		POLYGON_OFFSET_LINE:                             gl.POLYGON_OFFSET_LINE,
		POLYGON_OFFSET_POINT:                            gl.POLYGON_OFFSET_POINT,
		POLYGON_OFFSET_UNITS:                            gl.POLYGON_OFFSET_UNITS,
		POLYGON_SMOOTH:                                  gl.POLYGON_SMOOTH,
		POLYGON_SMOOTH_HINT:                             gl.POLYGON_SMOOTH_HINT,
		PRIMITIVES_GENERATED:                            gl.PRIMITIVES_GENERATED,
		PRIMITIVES_SUBMITTED:                            gl.PRIMITIVES_SUBMITTED,
		PRIMITIVES_SUBMITTED_ARB:                        gl.PRIMITIVES_SUBMITTED_ARB,
		PRIMITIVE_BOUNDING_BOX_ARB:                      gl.PRIMITIVE_BOUNDING_BOX_ARB,
		PRIMITIVE_RESTART:                               gl.PRIMITIVE_RESTART,
		PRIMITIVE_RESTART_FIXED_INDEX:                   gl.PRIMITIVE_RESTART_FIXED_INDEX,
		PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED:         gl.PRIMITIVE_RESTART_FOR_PATCHES_SUPPORTED,
		PRIMITIVE_RESTART_INDEX:                         gl.PRIMITIVE_RESTART_INDEX,
		PROGRAM:                                         gl.PROGRAM,
		PROGRAMMABLE_SAMPLE_LOCATION_ARB:                gl.PROGRAMMABLE_SAMPLE_LOCATION_ARB,
		PROGRAMMABLE_SAMPLE_LOCATION_NV:                 gl.PROGRAMMABLE_SAMPLE_LOCATION_NV,
		PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_ARB:     gl.PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_ARB,
		PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_NV:      gl.PROGRAMMABLE_SAMPLE_LOCATION_TABLE_SIZE_NV,
		PROGRAM_BINARY_FORMATS:                          gl.PROGRAM_BINARY_FORMATS,
		PROGRAM_BINARY_LENGTH:                           gl.PROGRAM_BINARY_LENGTH,
		PROGRAM_BINARY_RETRIEVABLE_HINT:                 gl.PROGRAM_BINARY_RETRIEVABLE_HINT,
		PROGRAM_INPUT:                                   gl.PROGRAM_INPUT,
		PROGRAM_KHR:                                     gl.PROGRAM_KHR,
		PROGRAM_MATRIX_EXT:                              gl.PROGRAM_MATRIX_EXT,
		PROGRAM_MATRIX_STACK_DEPTH_EXT:                  gl.PROGRAM_MATRIX_STACK_DEPTH_EXT,
		PROGRAM_OBJECT_EXT:                              gl.PROGRAM_OBJECT_EXT,
		PROGRAM_OUTPUT:                                  gl.PROGRAM_OUTPUT,
		PROGRAM_PIPELINE:                                gl.PROGRAM_PIPELINE,
		PROGRAM_PIPELINE_BINDING:                        gl.PROGRAM_PIPELINE_BINDING,
		PROGRAM_PIPELINE_BINDING_EXT:                    gl.PROGRAM_PIPELINE_BINDING_EXT,
		PROGRAM_PIPELINE_KHR:                            gl.PROGRAM_PIPELINE_KHR,
		PROGRAM_PIPELINE_OBJECT_EXT:                     gl.PROGRAM_PIPELINE_OBJECT_EXT,
		PROGRAM_POINT_SIZE:                              gl.PROGRAM_POINT_SIZE,
		PROGRAM_POINT_SIZE_ARB:                          gl.PROGRAM_POINT_SIZE_ARB,
		PROGRAM_SEPARABLE:                               gl.PROGRAM_SEPARABLE,
		PROGRAM_SEPARABLE_EXT:                           gl.PROGRAM_SEPARABLE_EXT,
		PROVOKING_VERTEX:                                gl.PROVOKING_VERTEX,
		PROXY_TEXTURE_1D:                                gl.PROXY_TEXTURE_1D,
		PROXY_TEXTURE_1D_ARRAY:                          gl.PROXY_TEXTURE_1D_ARRAY,
		PROXY_TEXTURE_2D:                                gl.PROXY_TEXTURE_2D,
		PROXY_TEXTURE_2D_ARRAY:                          gl.PROXY_TEXTURE_2D_ARRAY,
		PROXY_TEXTURE_2D_MULTISAMPLE:                    gl.PROXY_TEXTURE_2D_MULTISAMPLE,
		PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY:              gl.PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY,
		PROXY_TEXTURE_3D:                                gl.PROXY_TEXTURE_3D,
		PROXY_TEXTURE_CUBE_MAP:                          gl.PROXY_TEXTURE_CUBE_MAP,
		PROXY_TEXTURE_CUBE_MAP_ARRAY:                    gl.PROXY_TEXTURE_CUBE_MAP_ARRAY,
		PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB:                gl.PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB,
		PROXY_TEXTURE_RECTANGLE:                         gl.PROXY_TEXTURE_RECTANGLE,
		QUADRATIC_CURVE_TO_NV:                           gl.QUADRATIC_CURVE_TO_NV,
		QUADS:                                           gl.QUADS,
		QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION:        gl.QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION,
		QUERY:                                           gl.QUERY,
		QUERY_BUFFER:                                    gl.QUERY_BUFFER,
		QUERY_BUFFER_BARRIER_BIT:                        gl.QUERY_BUFFER_BARRIER_BIT,
		QUERY_BUFFER_BINDING:                            gl.QUERY_BUFFER_BINDING,
		QUERY_BY_REGION_NO_WAIT:                         gl.QUERY_BY_REGION_NO_WAIT,
		QUERY_BY_REGION_NO_WAIT_INVERTED:                gl.QUERY_BY_REGION_NO_WAIT_INVERTED,
		QUERY_BY_REGION_NO_WAIT_NV:                      gl.QUERY_BY_REGION_NO_WAIT_NV,
		QUERY_BY_REGION_WAIT:                            gl.QUERY_BY_REGION_WAIT,
		QUERY_BY_REGION_WAIT_INVERTED:                   gl.QUERY_BY_REGION_WAIT_INVERTED,
		QUERY_BY_REGION_WAIT_NV:                         gl.QUERY_BY_REGION_WAIT_NV,
		QUERY_COUNTER_BITS:                              gl.QUERY_COUNTER_BITS,
		QUERY_KHR:                                       gl.QUERY_KHR,
		QUERY_NO_WAIT:                                   gl.QUERY_NO_WAIT,
		QUERY_NO_WAIT_INVERTED:                          gl.QUERY_NO_WAIT_INVERTED,
		QUERY_NO_WAIT_NV:                                gl.QUERY_NO_WAIT_NV,
		QUERY_OBJECT_EXT:                                gl.QUERY_OBJECT_EXT,
		QUERY_RESULT:                                    gl.QUERY_RESULT,
		QUERY_RESULT_AVAILABLE:                          gl.QUERY_RESULT_AVAILABLE,
		QUERY_RESULT_NO_WAIT:                            gl.QUERY_RESULT_NO_WAIT,
		QUERY_TARGET:                                    gl.QUERY_TARGET,
		QUERY_WAIT:                                      gl.QUERY_WAIT,
		QUERY_WAIT_INVERTED:                             gl.QUERY_WAIT_INVERTED,
		QUERY_WAIT_NV:                                   gl.QUERY_WAIT_NV,
		R11F_G11F_B10F:                                  gl.R11F_G11F_B10F,
		R16:                                             gl.R16,
		R16F:                                            gl.R16F,
		R16I:                                            gl.R16I,
		R16UI:                                           gl.R16UI,
		R16_SNORM:                                       gl.R16_SNORM,
		R32F:                                            gl.R32F,
		R32I:                                            gl.R32I,
		R32UI:                                           gl.R32UI,
		R3_G3_B2:                                        gl.R3_G3_B2,
		R8:                                              gl.R8,
		R8I:                                             gl.R8I,
		R8UI:                                            gl.R8UI,
		R8_SNORM:                                        gl.R8_SNORM,
		RASTERIZER_DISCARD:                              gl.RASTERIZER_DISCARD,
		RASTER_FIXED_SAMPLE_LOCATIONS_EXT:               gl.RASTER_FIXED_SAMPLE_LOCATIONS_EXT,
		RASTER_MULTISAMPLE_EXT:                          gl.RASTER_MULTISAMPLE_EXT,
		RASTER_SAMPLES_EXT:                              gl.RASTER_SAMPLES_EXT,
		READ_BUFFER:                                     gl.READ_BUFFER,
		READ_FRAMEBUFFER:                                gl.READ_FRAMEBUFFER,
		READ_FRAMEBUFFER_BINDING:                        gl.READ_FRAMEBUFFER_BINDING,
		READ_ONLY:                                       gl.READ_ONLY,
		READ_PIXELS:                                     gl.READ_PIXELS,
		READ_PIXELS_FORMAT:                              gl.READ_PIXELS_FORMAT,
		READ_PIXELS_TYPE:                                gl.READ_PIXELS_TYPE,
		READ_WRITE:                                      gl.READ_WRITE,
		RECT_NV:                                         gl.RECT_NV,
		RED:                                             gl.RED,
		RED_INTEGER:                                     gl.RED_INTEGER,
		RED_NV:                                          gl.RED_NV,
		REFERENCED_BY_COMPUTE_SHADER:                    gl.REFERENCED_BY_COMPUTE_SHADER,
		REFERENCED_BY_FRAGMENT_SHADER:                   gl.REFERENCED_BY_FRAGMENT_SHADER,
		REFERENCED_BY_GEOMETRY_SHADER:                   gl.REFERENCED_BY_GEOMETRY_SHADER,
		REFERENCED_BY_TESS_CONTROL_SHADER:               gl.REFERENCED_BY_TESS_CONTROL_SHADER,
		REFERENCED_BY_TESS_EVALUATION_SHADER:            gl.REFERENCED_BY_TESS_EVALUATION_SHADER,
		REFERENCED_BY_VERTEX_SHADER:                     gl.REFERENCED_BY_VERTEX_SHADER,
		RELATIVE_ARC_TO_NV:                              gl.RELATIVE_ARC_TO_NV,
		RELATIVE_CONIC_CURVE_TO_NV:                      gl.RELATIVE_CONIC_CURVE_TO_NV,
		RELATIVE_CUBIC_CURVE_TO_NV:                      gl.RELATIVE_CUBIC_CURVE_TO_NV,
		RELATIVE_HORIZONTAL_LINE_TO_NV:                  gl.RELATIVE_HORIZONTAL_LINE_TO_NV,
		RELATIVE_LARGE_CCW_ARC_TO_NV:                    gl.RELATIVE_LARGE_CCW_ARC_TO_NV,
		RELATIVE_LARGE_CW_ARC_TO_NV:                     gl.RELATIVE_LARGE_CW_ARC_TO_NV,
		RELATIVE_LINE_TO_NV:                             gl.RELATIVE_LINE_TO_NV,
		RELATIVE_MOVE_TO_NV:                             gl.RELATIVE_MOVE_TO_NV,
		RELATIVE_QUADRATIC_CURVE_TO_NV:                  gl.RELATIVE_QUADRATIC_CURVE_TO_NV,
		RELATIVE_RECT_NV:                                gl.RELATIVE_RECT_NV,
		RELATIVE_ROUNDED_RECT2_NV:                       gl.RELATIVE_ROUNDED_RECT2_NV,
		RELATIVE_ROUNDED_RECT4_NV:                       gl.RELATIVE_ROUNDED_RECT4_NV,
		RELATIVE_ROUNDED_RECT8_NV:                       gl.RELATIVE_ROUNDED_RECT8_NV,
		RELATIVE_ROUNDED_RECT_NV:                        gl.RELATIVE_ROUNDED_RECT_NV,
		RELATIVE_SMALL_CCW_ARC_TO_NV:                    gl.RELATIVE_SMALL_CCW_ARC_TO_NV,
		RELATIVE_SMALL_CW_ARC_TO_NV:                     gl.RELATIVE_SMALL_CW_ARC_TO_NV,
		RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV:               gl.RELATIVE_SMOOTH_CUBIC_CURVE_TO_NV,
		RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV:           gl.RELATIVE_SMOOTH_QUADRATIC_CURVE_TO_NV,
		RELATIVE_VERTICAL_LINE_TO_NV:                    gl.RELATIVE_VERTICAL_LINE_TO_NV,
		RENDERBUFFER:                                    gl.RENDERBUFFER,
		RENDERBUFFER_ALPHA_SIZE:                         gl.RENDERBUFFER_ALPHA_SIZE,
		RENDERBUFFER_BINDING:                            gl.RENDERBUFFER_BINDING,
		RENDERBUFFER_BLUE_SIZE:                          gl.RENDERBUFFER_BLUE_SIZE,
		RENDERBUFFER_COLOR_SAMPLES_NV:                   gl.RENDERBUFFER_COLOR_SAMPLES_NV,
		RENDERBUFFER_COVERAGE_SAMPLES_NV:                gl.RENDERBUFFER_COVERAGE_SAMPLES_NV,
		RENDERBUFFER_DEPTH_SIZE:                         gl.RENDERBUFFER_DEPTH_SIZE,
		RENDERBUFFER_GREEN_SIZE:                         gl.RENDERBUFFER_GREEN_SIZE,
		RENDERBUFFER_HEIGHT:                             gl.RENDERBUFFER_HEIGHT,
		RENDERBUFFER_INTERNAL_FORMAT:                    gl.RENDERBUFFER_INTERNAL_FORMAT,
		RENDERBUFFER_RED_SIZE:                           gl.RENDERBUFFER_RED_SIZE,
		RENDERBUFFER_SAMPLES:                            gl.RENDERBUFFER_SAMPLES,
		RENDERBUFFER_STENCIL_SIZE:                       gl.RENDERBUFFER_STENCIL_SIZE,
		RENDERBUFFER_WIDTH:                              gl.RENDERBUFFER_WIDTH,
		RENDERER:                                        gl.RENDERER,
		REPEAT:                                          gl.REPEAT,
		REPLACE:                                         gl.REPLACE,
		RESET_NOTIFICATION_STRATEGY:                     gl.RESET_NOTIFICATION_STRATEGY,
		RESET_NOTIFICATION_STRATEGY_ARB:                 gl.RESET_NOTIFICATION_STRATEGY_ARB,
		RESET_NOTIFICATION_STRATEGY_KHR:                 gl.RESET_NOTIFICATION_STRATEGY_KHR,
		RESTART_PATH_NV:                                 gl.RESTART_PATH_NV,
		RG:                                              gl.RG,
		RG16:                                            gl.RG16,
		RG16F:                                           gl.RG16F,
		RG16I:                                           gl.RG16I,
		RG16UI:                                          gl.RG16UI,
		RG16_SNORM:                                      gl.RG16_SNORM,
		RG32F:                                           gl.RG32F,
		RG32I:                                           gl.RG32I,
		RG32UI:                                          gl.RG32UI,
		RG8:                                             gl.RG8,
		RG8I:                                            gl.RG8I,
		RG8UI:                                           gl.RG8UI,
		RG8_SNORM:                                       gl.RG8_SNORM,
		RGB:                                             gl.RGB,
		RGB10:                                           gl.RGB10,
		RGB10_A2:                                        gl.RGB10_A2,
		RGB10_A2UI:                                      gl.RGB10_A2UI,
		RGB12:                                           gl.RGB12,
		RGB16:                                           gl.RGB16,
		RGB16F:                                          gl.RGB16F,
		RGB16I:                                          gl.RGB16I,
		RGB16UI:                                         gl.RGB16UI,
		RGB16_SNORM:                                     gl.RGB16_SNORM,
		RGB32F:                                          gl.RGB32F,
		RGB32I:                                          gl.RGB32I,
		RGB32UI:                                         gl.RGB32UI,
		RGB4:                                            gl.RGB4,
		RGB5:                                            gl.RGB5,
		RGB565:                                          gl.RGB565,
		RGB5_A1:                                         gl.RGB5_A1,
		RGB8:                                            gl.RGB8,
		RGB8I:                                           gl.RGB8I,
		RGB8UI:                                          gl.RGB8UI,
		RGB8_SNORM:                                      gl.RGB8_SNORM,
		RGB9_E5:                                         gl.RGB9_E5,
		RGBA:                                            gl.RGBA,
		RGBA12:                                          gl.RGBA12,
		RGBA16:                                          gl.RGBA16,
		RGBA16F:                                         gl.RGBA16F,
		RGBA16I:                                         gl.RGBA16I,
		RGBA16UI:                                        gl.RGBA16UI,
		RGBA16_SNORM:                                    gl.RGBA16_SNORM,
		RGBA2:                                           gl.RGBA2,
		RGBA32F:                                         gl.RGBA32F,
		RGBA32I:                                         gl.RGBA32I,
		RGBA32UI:                                        gl.RGBA32UI,
		RGBA4:                                           gl.RGBA4,
		RGBA8:                                           gl.RGBA8,
		RGBA8I:                                          gl.RGBA8I,
		RGBA8UI:                                         gl.RGBA8UI,
		RGBA8_SNORM:                                     gl.RGBA8_SNORM,
		RGBA_INTEGER:                                    gl.RGBA_INTEGER,
		RGB_422_APPLE:                                   gl.RGB_422_APPLE,
		RGB_INTEGER:                                     gl.RGB_INTEGER,
		RGB_RAW_422_APPLE:                               gl.RGB_RAW_422_APPLE,
		RG_INTEGER:                                      gl.RG_INTEGER,
		RIGHT:                                           gl.RIGHT,
		ROUNDED_RECT2_NV:                                gl.ROUNDED_RECT2_NV,
		ROUNDED_RECT4_NV:                                gl.ROUNDED_RECT4_NV,
		ROUNDED_RECT8_NV:                                gl.ROUNDED_RECT8_NV,
		ROUNDED_RECT_NV:                                 gl.ROUNDED_RECT_NV,
		ROUND_NV:                                        gl.ROUND_NV,
		SAMPLER:                                         gl.SAMPLER,
		SAMPLER_1D:                                      gl.SAMPLER_1D,
		SAMPLER_1D_ARRAY:                                gl.SAMPLER_1D_ARRAY,
		SAMPLER_1D_ARRAY_SHADOW:                         gl.SAMPLER_1D_ARRAY_SHADOW,
		SAMPLER_1D_SHADOW:                               gl.SAMPLER_1D_SHADOW,
		SAMPLER_2D:                                      gl.SAMPLER_2D,
		SAMPLER_2D_ARRAY:                                gl.SAMPLER_2D_ARRAY,
		SAMPLER_2D_ARRAY_SHADOW:                         gl.SAMPLER_2D_ARRAY_SHADOW,
		SAMPLER_2D_MULTISAMPLE:                          gl.SAMPLER_2D_MULTISAMPLE,
		SAMPLER_2D_MULTISAMPLE_ARRAY:                    gl.SAMPLER_2D_MULTISAMPLE_ARRAY,
		SAMPLER_2D_RECT:                                 gl.SAMPLER_2D_RECT,
		SAMPLER_2D_RECT_SHADOW:                          gl.SAMPLER_2D_RECT_SHADOW,
		SAMPLER_2D_SHADOW:                               gl.SAMPLER_2D_SHADOW,
		SAMPLER_3D:                                      gl.SAMPLER_3D,
		SAMPLER_BINDING:                                 gl.SAMPLER_BINDING,
		SAMPLER_BUFFER:                                  gl.SAMPLER_BUFFER,
		SAMPLER_CUBE:                                    gl.SAMPLER_CUBE,
		SAMPLER_CUBE_MAP_ARRAY:                          gl.SAMPLER_CUBE_MAP_ARRAY,
		SAMPLER_CUBE_MAP_ARRAY_ARB:                      gl.SAMPLER_CUBE_MAP_ARRAY_ARB,
		SAMPLER_CUBE_MAP_ARRAY_SHADOW:                   gl.SAMPLER_CUBE_MAP_ARRAY_SHADOW,
		SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB:               gl.SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB,
		SAMPLER_CUBE_SHADOW:                             gl.SAMPLER_CUBE_SHADOW,
		SAMPLER_KHR:                                     gl.SAMPLER_KHR,
		SAMPLES:                                         gl.SAMPLES,
		SAMPLES_PASSED:                                  gl.SAMPLES_PASSED,
		SAMPLE_ALPHA_TO_COVERAGE:                        gl.SAMPLE_ALPHA_TO_COVERAGE,
		SAMPLE_ALPHA_TO_ONE:                             gl.SAMPLE_ALPHA_TO_ONE,
		SAMPLE_BUFFERS:                                  gl.SAMPLE_BUFFERS,
		SAMPLE_COVERAGE:                                 gl.SAMPLE_COVERAGE,
		SAMPLE_COVERAGE_INVERT:                          gl.SAMPLE_COVERAGE_INVERT,
		SAMPLE_COVERAGE_VALUE:                           gl.SAMPLE_COVERAGE_VALUE,
		SAMPLE_LOCATION_ARB:                             gl.SAMPLE_LOCATION_ARB,
		SAMPLE_LOCATION_NV:                              gl.SAMPLE_LOCATION_NV,
		SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_ARB:           gl.SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_ARB,
		SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_NV:            gl.SAMPLE_LOCATION_PIXEL_GRID_HEIGHT_NV,
		SAMPLE_LOCATION_PIXEL_GRID_WIDTH_ARB:            gl.SAMPLE_LOCATION_PIXEL_GRID_WIDTH_ARB,
		SAMPLE_LOCATION_PIXEL_GRID_WIDTH_NV:             gl.SAMPLE_LOCATION_PIXEL_GRID_WIDTH_NV,
		SAMPLE_LOCATION_SUBPIXEL_BITS_ARB:               gl.SAMPLE_LOCATION_SUBPIXEL_BITS_ARB,
		SAMPLE_LOCATION_SUBPIXEL_BITS_NV:                gl.SAMPLE_LOCATION_SUBPIXEL_BITS_NV,
		SAMPLE_MASK:                                     gl.SAMPLE_MASK,
		SAMPLE_MASK_VALUE:                               gl.SAMPLE_MASK_VALUE,
		SAMPLE_POSITION:                                 gl.SAMPLE_POSITION,
		SAMPLE_SHADING:                                  gl.SAMPLE_SHADING,
		SAMPLE_SHADING_ARB:                              gl.SAMPLE_SHADING_ARB,
		SCISSOR_BOX:                                     gl.SCISSOR_BOX,
		SCISSOR_COMMAND_NV:                              gl.SCISSOR_COMMAND_NV,
		SCISSOR_TEST:                                    gl.SCISSOR_TEST,
		SCREEN_KHR:                                      gl.SCREEN_KHR,
		SCREEN_NV:                                       gl.SCREEN_NV,
		SECONDARY_COLOR_ARRAY_ADDRESS_NV:                gl.SECONDARY_COLOR_ARRAY_ADDRESS_NV,
		SECONDARY_COLOR_ARRAY_LENGTH_NV:                 gl.SECONDARY_COLOR_ARRAY_LENGTH_NV,
		SEPARATE_ATTRIBS:                                gl.SEPARATE_ATTRIBS,
		SET:                                             gl.SET,
		SHADER:                                          gl.SHADER,
		SHADER_BINARY_FORMATS:                           gl.SHADER_BINARY_FORMATS,
		SHADER_BINARY_FORMAT_SPIR_V:                     gl.SHADER_BINARY_FORMAT_SPIR_V,
		SHADER_BINARY_FORMAT_SPIR_V_ARB:                 gl.SHADER_BINARY_FORMAT_SPIR_V_ARB,
		SHADER_COMPILER:                                 gl.SHADER_COMPILER,
		SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV:             gl.SHADER_GLOBAL_ACCESS_BARRIER_BIT_NV,
		SHADER_IMAGE_ACCESS_BARRIER_BIT:                 gl.SHADER_IMAGE_ACCESS_BARRIER_BIT,
		SHADER_IMAGE_ATOMIC:                             gl.SHADER_IMAGE_ATOMIC,
		SHADER_IMAGE_LOAD:                               gl.SHADER_IMAGE_LOAD,
		SHADER_IMAGE_STORE:                              gl.SHADER_IMAGE_STORE,
		SHADER_INCLUDE_ARB:                              gl.SHADER_INCLUDE_ARB,
		SHADER_KHR:                                      gl.SHADER_KHR,
		SHADER_OBJECT_EXT:                               gl.SHADER_OBJECT_EXT,
		SHADER_SOURCE_LENGTH:                            gl.SHADER_SOURCE_LENGTH,
		SHADER_STORAGE_BARRIER_BIT:                      gl.SHADER_STORAGE_BARRIER_BIT,
		SHADER_STORAGE_BLOCK:                            gl.SHADER_STORAGE_BLOCK,
		SHADER_STORAGE_BUFFER:                           gl.SHADER_STORAGE_BUFFER,
		SHADER_STORAGE_BUFFER_BINDING:                   gl.SHADER_STORAGE_BUFFER_BINDING,
		SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT:          gl.SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT,
		SHADER_STORAGE_BUFFER_SIZE:                      gl.SHADER_STORAGE_BUFFER_SIZE,
		SHADER_STORAGE_BUFFER_START:                     gl.SHADER_STORAGE_BUFFER_START,
		SHADER_TYPE:                                     gl.SHADER_TYPE,
		SHADING_LANGUAGE_VERSION:                        gl.SHADING_LANGUAGE_VERSION,
		SHARED_EDGE_NV:                                  gl.SHARED_EDGE_NV,
		SHORT:                                           gl.SHORT,
		SIGNALED:                                        gl.SIGNALED,
		SIGNED_NORMALIZED:                               gl.SIGNED_NORMALIZED,
		SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST:             gl.SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST,
		SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE:            gl.SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE,
		SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST:           gl.SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST,
		SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE:          gl.SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE,
		SKIP_DECODE_EXT:                                 gl.SKIP_DECODE_EXT,
		SKIP_MISSING_GLYPH_NV:                           gl.SKIP_MISSING_GLYPH_NV,
		SMALL_CCW_ARC_TO_NV:                             gl.SMALL_CCW_ARC_TO_NV,
		SMALL_CW_ARC_TO_NV:                              gl.SMALL_CW_ARC_TO_NV,
		SMOOTH_CUBIC_CURVE_TO_NV:                        gl.SMOOTH_CUBIC_CURVE_TO_NV,
		SMOOTH_LINE_WIDTH_GRANULARITY:                   gl.SMOOTH_LINE_WIDTH_GRANULARITY,
		SMOOTH_LINE_WIDTH_RANGE:                         gl.SMOOTH_LINE_WIDTH_RANGE,
		SMOOTH_POINT_SIZE_GRANULARITY:                   gl.SMOOTH_POINT_SIZE_GRANULARITY,
		SMOOTH_POINT_SIZE_RANGE:                         gl.SMOOTH_POINT_SIZE_RANGE,
		SMOOTH_QUADRATIC_CURVE_TO_NV:                    gl.SMOOTH_QUADRATIC_CURVE_TO_NV,
		SM_COUNT_NV:                                     gl.SM_COUNT_NV,
		SOFTLIGHT_KHR:                                   gl.SOFTLIGHT_KHR,
		SOFTLIGHT_NV:                                    gl.SOFTLIGHT_NV,
		SPARSE_BUFFER_PAGE_SIZE_ARB:                     gl.SPARSE_BUFFER_PAGE_SIZE_ARB,
		SPARSE_STORAGE_BIT_ARB:                          gl.SPARSE_STORAGE_BIT_ARB,
		SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB:      gl.SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB,
		SPIR_V_BINARY:                                   gl.SPIR_V_BINARY,
		SPIR_V_BINARY_ARB:                               gl.SPIR_V_BINARY_ARB,
		SPIR_V_EXTENSIONS:                               gl.SPIR_V_EXTENSIONS,
		SQUARE_NV:                                       gl.SQUARE_NV,
		SRC1_ALPHA:                                      gl.SRC1_ALPHA,
		SRC1_COLOR:                                      gl.SRC1_COLOR,
		SRC_ALPHA:                                       gl.SRC_ALPHA,
		SRC_ALPHA_SATURATE:                              gl.SRC_ALPHA_SATURATE,
		SRC_ATOP_NV:                                     gl.SRC_ATOP_NV,
		SRC_COLOR:                                       gl.SRC_COLOR,
		SRC_IN_NV:                                       gl.SRC_IN_NV,
		SRC_NV:                                          gl.SRC_NV,
		SRC_OUT_NV:                                      gl.SRC_OUT_NV,
		SRC_OVER_NV:                                     gl.SRC_OVER_NV,
		SRGB:                                            gl.SRGB,
		SRGB8:                                           gl.SRGB8,
		SRGB8_ALPHA8:                                    gl.SRGB8_ALPHA8,
		SRGB_ALPHA:                                      gl.SRGB_ALPHA,
		SRGB_DECODE_ARB:                                 gl.SRGB_DECODE_ARB,
		SRGB_READ:                                       gl.SRGB_READ,
		SRGB_WRITE:                                      gl.SRGB_WRITE,
		STACK_OVERFLOW:                                  gl.STACK_OVERFLOW,
		STACK_OVERFLOW_KHR:                              gl.STACK_OVERFLOW_KHR,
		STACK_UNDERFLOW:                                 gl.STACK_UNDERFLOW,
		STACK_UNDERFLOW_KHR:                             gl.STACK_UNDERFLOW_KHR,
		STANDARD_FONT_FORMAT_NV:                         gl.STANDARD_FONT_FORMAT_NV,
		STANDARD_FONT_NAME_NV:                           gl.STANDARD_FONT_NAME_NV,
		STATIC_COPY:                                     gl.STATIC_COPY,
		STATIC_DRAW:                                     gl.STATIC_DRAW,
		STATIC_READ:                                     gl.STATIC_READ,
		STENCIL:                                         gl.STENCIL,
		STENCIL_ATTACHMENT:                              gl.STENCIL_ATTACHMENT,
		STENCIL_BACK_FAIL:                               gl.STENCIL_BACK_FAIL,
		STENCIL_BACK_FUNC:                               gl.STENCIL_BACK_FUNC,
		STENCIL_BACK_PASS_DEPTH_FAIL:                    gl.STENCIL_BACK_PASS_DEPTH_FAIL,
		STENCIL_BACK_PASS_DEPTH_PASS:                    gl.STENCIL_BACK_PASS_DEPTH_PASS,
		STENCIL_BACK_REF:                                gl.STENCIL_BACK_REF,
		STENCIL_BACK_VALUE_MASK:                         gl.STENCIL_BACK_VALUE_MASK,
		STENCIL_BACK_WRITEMASK:                          gl.STENCIL_BACK_WRITEMASK,
		STENCIL_BUFFER_BIT:                              gl.STENCIL_BUFFER_BIT,
		STENCIL_CLEAR_VALUE:                             gl.STENCIL_CLEAR_VALUE,
		STENCIL_COMPONENTS:                              gl.STENCIL_COMPONENTS,
		STENCIL_FAIL:                                    gl.STENCIL_FAIL,
		STENCIL_FUNC:                                    gl.STENCIL_FUNC,
		STENCIL_INDEX:                                   gl.STENCIL_INDEX,
		STENCIL_INDEX1:                                  gl.STENCIL_INDEX1,
		STENCIL_INDEX16:                                 gl.STENCIL_INDEX16,
		STENCIL_INDEX4:                                  gl.STENCIL_INDEX4,
		STENCIL_INDEX8:                                  gl.STENCIL_INDEX8,
		STENCIL_PASS_DEPTH_FAIL:                         gl.STENCIL_PASS_DEPTH_FAIL,
		STENCIL_PASS_DEPTH_PASS:                         gl.STENCIL_PASS_DEPTH_PASS,
		STENCIL_REF:                                     gl.STENCIL_REF,
		STENCIL_REF_COMMAND_NV:                          gl.STENCIL_REF_COMMAND_NV,
		STENCIL_RENDERABLE:                              gl.STENCIL_RENDERABLE,
		STENCIL_SAMPLES_NV:                              gl.STENCIL_SAMPLES_NV,
		STENCIL_TEST:                                    gl.STENCIL_TEST,
		STENCIL_VALUE_MASK:                              gl.STENCIL_VALUE_MASK,
		STENCIL_WRITEMASK:                               gl.STENCIL_WRITEMASK,
		STEREO:                                          gl.STEREO,
		STREAM_COPY:                                     gl.STREAM_COPY,
		STREAM_DRAW:                                     gl.STREAM_DRAW,
		STREAM_READ:                                     gl.STREAM_READ,
		SUBPIXEL_BITS:                                   gl.SUBPIXEL_BITS,
		SUBPIXEL_PRECISION_BIAS_X_BITS_NV:               gl.SUBPIXEL_PRECISION_BIAS_X_BITS_NV,
		SUBPIXEL_PRECISION_BIAS_Y_BITS_NV:               gl.SUBPIXEL_PRECISION_BIAS_Y_BITS_NV,
		SUPERSAMPLE_SCALE_X_NV:                          gl.SUPERSAMPLE_SCALE_X_NV,
		SUPERSAMPLE_SCALE_Y_NV:                          gl.SUPERSAMPLE_SCALE_Y_NV,
		SYNC_CL_EVENT_ARB:                               gl.SYNC_CL_EVENT_ARB,
		SYNC_CL_EVENT_COMPLETE_ARB:                      gl.SYNC_CL_EVENT_COMPLETE_ARB,
		SYNC_CONDITION:                                  gl.SYNC_CONDITION,
		SYNC_FENCE:                                      gl.SYNC_FENCE,
		SYNC_FLAGS:                                      gl.SYNC_FLAGS,
		SYNC_FLUSH_COMMANDS_BIT:                         gl.SYNC_FLUSH_COMMANDS_BIT,
		SYNC_GPU_COMMANDS_COMPLETE:                      gl.SYNC_GPU_COMMANDS_COMPLETE,
		SYNC_STATUS:                                     gl.SYNC_STATUS,
		SYSTEM_FONT_NAME_NV:                             gl.SYSTEM_FONT_NAME_NV,
		TERMINATE_SEQUENCE_COMMAND_NV:                   gl.TERMINATE_SEQUENCE_COMMAND_NV,
		TESS_CONTROL_OUTPUT_VERTICES:                    gl.TESS_CONTROL_OUTPUT_VERTICES,
		TESS_CONTROL_SHADER:                             gl.TESS_CONTROL_SHADER,
		TESS_CONTROL_SHADER_BIT:                         gl.TESS_CONTROL_SHADER_BIT,
		TESS_CONTROL_SHADER_PATCHES:                     gl.TESS_CONTROL_SHADER_PATCHES,
		TESS_CONTROL_SHADER_PATCHES_ARB:                 gl.TESS_CONTROL_SHADER_PATCHES_ARB,
		TESS_CONTROL_SUBROUTINE:                         gl.TESS_CONTROL_SUBROUTINE,
		TESS_CONTROL_SUBROUTINE_UNIFORM:                 gl.TESS_CONTROL_SUBROUTINE_UNIFORM,
		TESS_CONTROL_TEXTURE:                            gl.TESS_CONTROL_TEXTURE,
		TESS_EVALUATION_SHADER:                          gl.TESS_EVALUATION_SHADER,
		TESS_EVALUATION_SHADER_BIT:                      gl.TESS_EVALUATION_SHADER_BIT,
		TESS_EVALUATION_SHADER_INVOCATIONS:              gl.TESS_EVALUATION_SHADER_INVOCATIONS,
		TESS_EVALUATION_SHADER_INVOCATIONS_ARB:          gl.TESS_EVALUATION_SHADER_INVOCATIONS_ARB,
		TESS_EVALUATION_SUBROUTINE:                      gl.TESS_EVALUATION_SUBROUTINE,
		TESS_EVALUATION_SUBROUTINE_UNIFORM:              gl.TESS_EVALUATION_SUBROUTINE_UNIFORM,
		TESS_EVALUATION_TEXTURE:                         gl.TESS_EVALUATION_TEXTURE,
		TESS_GEN_MODE:                                   gl.TESS_GEN_MODE,
		TESS_GEN_POINT_MODE:                             gl.TESS_GEN_POINT_MODE,
		TESS_GEN_SPACING:                                gl.TESS_GEN_SPACING,
		TESS_GEN_VERTEX_ORDER:                           gl.TESS_GEN_VERTEX_ORDER,
		TEXTURE:                                         gl.TEXTURE,
		TEXTURE0:                                        gl.TEXTURE0,
		TEXTURE1:                                        gl.TEXTURE1,
		TEXTURE10:                                       gl.TEXTURE10,
		TEXTURE11:                                       gl.TEXTURE11,
		TEXTURE12:                                       gl.TEXTURE12,
		TEXTURE13:                                       gl.TEXTURE13,
		TEXTURE14:                                       gl.TEXTURE14,
		TEXTURE15:                                       gl.TEXTURE15,
		TEXTURE16:                                       gl.TEXTURE16,
		TEXTURE17:                                       gl.TEXTURE17,
		TEXTURE18:                                       gl.TEXTURE18,
		TEXTURE19:                                       gl.TEXTURE19,
		TEXTURE2:                                        gl.TEXTURE2,
		TEXTURE20:                                       gl.TEXTURE20,
		TEXTURE21:                                       gl.TEXTURE21,
		TEXTURE22:                                       gl.TEXTURE22,
		TEXTURE23:                                       gl.TEXTURE23,
		TEXTURE24:                                       gl.TEXTURE24,
		TEXTURE25:                                       gl.TEXTURE25,
		TEXTURE26:                                       gl.TEXTURE26,
		TEXTURE27:                                       gl.TEXTURE27,
		TEXTURE28:                                       gl.TEXTURE28,
		TEXTURE29:                                       gl.TEXTURE29,
		TEXTURE3:                                        gl.TEXTURE3,
		TEXTURE30:                                       gl.TEXTURE30,
		TEXTURE31:                                       gl.TEXTURE31,
		TEXTURE4:                                        gl.TEXTURE4,
		TEXTURE5:                                        gl.TEXTURE5,
		TEXTURE6:                                        gl.TEXTURE6,
		TEXTURE7:                                        gl.TEXTURE7,
		TEXTURE8:                                        gl.TEXTURE8,
		TEXTURE9:                                        gl.TEXTURE9,
		TEXTURE_1D:                                      gl.TEXTURE_1D,
		TEXTURE_1D_ARRAY:                                gl.TEXTURE_1D_ARRAY,
		TEXTURE_2D:                                      gl.TEXTURE_2D,
		TEXTURE_2D_ARRAY:                                gl.TEXTURE_2D_ARRAY,
		TEXTURE_2D_MULTISAMPLE:                          gl.TEXTURE_2D_MULTISAMPLE,
		TEXTURE_2D_MULTISAMPLE_ARRAY:                    gl.TEXTURE_2D_MULTISAMPLE_ARRAY,
		TEXTURE_3D:                                      gl.TEXTURE_3D,
		TEXTURE_ALPHA_SIZE:                              gl.TEXTURE_ALPHA_SIZE,
		TEXTURE_ALPHA_TYPE:                              gl.TEXTURE_ALPHA_TYPE,
		TEXTURE_BASE_LEVEL:                              gl.TEXTURE_BASE_LEVEL,
		TEXTURE_BINDING_1D:                              gl.TEXTURE_BINDING_1D,
		TEXTURE_BINDING_1D_ARRAY:                        gl.TEXTURE_BINDING_1D_ARRAY,
		TEXTURE_BINDING_2D:                              gl.TEXTURE_BINDING_2D,
		TEXTURE_BINDING_2D_ARRAY:                        gl.TEXTURE_BINDING_2D_ARRAY,
		TEXTURE_BINDING_2D_MULTISAMPLE:                  gl.TEXTURE_BINDING_2D_MULTISAMPLE,
		TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY:            gl.TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY,
		TEXTURE_BINDING_3D:                              gl.TEXTURE_BINDING_3D,
		TEXTURE_BINDING_BUFFER:                          gl.TEXTURE_BINDING_BUFFER,
		TEXTURE_BINDING_BUFFER_ARB:                      gl.TEXTURE_BINDING_BUFFER_ARB,
		TEXTURE_BINDING_CUBE_MAP:                        gl.TEXTURE_BINDING_CUBE_MAP,
		TEXTURE_BINDING_CUBE_MAP_ARRAY:                  gl.TEXTURE_BINDING_CUBE_MAP_ARRAY,
		TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB:              gl.TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB,
		TEXTURE_BINDING_RECTANGLE:                       gl.TEXTURE_BINDING_RECTANGLE,
		TEXTURE_BLUE_SIZE:                               gl.TEXTURE_BLUE_SIZE,
		TEXTURE_BLUE_TYPE:                               gl.TEXTURE_BLUE_TYPE,
		TEXTURE_BORDER_COLOR:                            gl.TEXTURE_BORDER_COLOR,
		TEXTURE_BUFFER:                                  gl.TEXTURE_BUFFER,
		TEXTURE_BUFFER_ARB:                              gl.TEXTURE_BUFFER_ARB,
		TEXTURE_BUFFER_BINDING:                          gl.TEXTURE_BUFFER_BINDING,
		TEXTURE_BUFFER_DATA_STORE_BINDING:               gl.TEXTURE_BUFFER_DATA_STORE_BINDING,
		TEXTURE_BUFFER_DATA_STORE_BINDING_ARB:           gl.TEXTURE_BUFFER_DATA_STORE_BINDING_ARB,
		TEXTURE_BUFFER_FORMAT_ARB:                       gl.TEXTURE_BUFFER_FORMAT_ARB,
		TEXTURE_BUFFER_OFFSET:                           gl.TEXTURE_BUFFER_OFFSET,
		TEXTURE_BUFFER_OFFSET_ALIGNMENT:                 gl.TEXTURE_BUFFER_OFFSET_ALIGNMENT,
		TEXTURE_BUFFER_SIZE:                             gl.TEXTURE_BUFFER_SIZE,
		TEXTURE_COMPARE_FUNC:                            gl.TEXTURE_COMPARE_FUNC,
		TEXTURE_COMPARE_MODE:                            gl.TEXTURE_COMPARE_MODE,
		TEXTURE_COMPRESSED:                              gl.TEXTURE_COMPRESSED,
		TEXTURE_COMPRESSED_BLOCK_HEIGHT:                 gl.TEXTURE_COMPRESSED_BLOCK_HEIGHT,
		TEXTURE_COMPRESSED_BLOCK_SIZE:                   gl.TEXTURE_COMPRESSED_BLOCK_SIZE,
		TEXTURE_COMPRESSED_BLOCK_WIDTH:                  gl.TEXTURE_COMPRESSED_BLOCK_WIDTH,
		TEXTURE_COMPRESSED_IMAGE_SIZE:                   gl.TEXTURE_COMPRESSED_IMAGE_SIZE,
		TEXTURE_COMPRESSION_HINT:                        gl.TEXTURE_COMPRESSION_HINT,
		TEXTURE_COORD_ARRAY_ADDRESS_NV:                  gl.TEXTURE_COORD_ARRAY_ADDRESS_NV,
		TEXTURE_COORD_ARRAY_LENGTH_NV:                   gl.TEXTURE_COORD_ARRAY_LENGTH_NV,
		TEXTURE_CUBE_MAP:                                gl.TEXTURE_CUBE_MAP,
		TEXTURE_CUBE_MAP_ARRAY:                          gl.TEXTURE_CUBE_MAP_ARRAY,
		TEXTURE_CUBE_MAP_ARRAY_ARB:                      gl.TEXTURE_CUBE_MAP_ARRAY_ARB,
		TEXTURE_CUBE_MAP_NEGATIVE_X:                     gl.TEXTURE_CUBE_MAP_NEGATIVE_X,
		TEXTURE_CUBE_MAP_NEGATIVE_Y:                     gl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
		TEXTURE_CUBE_MAP_NEGATIVE_Z:                     gl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
		TEXTURE_CUBE_MAP_POSITIVE_X:                     gl.TEXTURE_CUBE_MAP_POSITIVE_X,
		TEXTURE_CUBE_MAP_POSITIVE_Y:                     gl.TEXTURE_CUBE_MAP_POSITIVE_Y,
		TEXTURE_CUBE_MAP_POSITIVE_Z:                     gl.TEXTURE_CUBE_MAP_POSITIVE_Z,
		TEXTURE_CUBE_MAP_SEAMLESS:                       gl.TEXTURE_CUBE_MAP_SEAMLESS,
		TEXTURE_DEPTH:                                   gl.TEXTURE_DEPTH,
		TEXTURE_DEPTH_SIZE:                              gl.TEXTURE_DEPTH_SIZE,
		TEXTURE_DEPTH_TYPE:                              gl.TEXTURE_DEPTH_TYPE,
		TEXTURE_FETCH_BARRIER_BIT:                       gl.TEXTURE_FETCH_BARRIER_BIT,
		TEXTURE_FIXED_SAMPLE_LOCATIONS:                  gl.TEXTURE_FIXED_SAMPLE_LOCATIONS,
		TEXTURE_GATHER:                                  gl.TEXTURE_GATHER,
		TEXTURE_GATHER_SHADOW:                           gl.TEXTURE_GATHER_SHADOW,
		TEXTURE_GREEN_SIZE:                              gl.TEXTURE_GREEN_SIZE,
		TEXTURE_GREEN_TYPE:                              gl.TEXTURE_GREEN_TYPE,
		TEXTURE_HEIGHT:                                  gl.TEXTURE_HEIGHT,
		TEXTURE_IMAGE_FORMAT:                            gl.TEXTURE_IMAGE_FORMAT,
		TEXTURE_IMAGE_TYPE:                              gl.TEXTURE_IMAGE_TYPE,
		TEXTURE_IMMUTABLE_FORMAT:                        gl.TEXTURE_IMMUTABLE_FORMAT,
		TEXTURE_IMMUTABLE_LEVELS:                        gl.TEXTURE_IMMUTABLE_LEVELS,
		TEXTURE_INTERNAL_FORMAT:                         gl.TEXTURE_INTERNAL_FORMAT,
		TEXTURE_LOD_BIAS:                                gl.TEXTURE_LOD_BIAS,
		TEXTURE_MAG_FILTER:                              gl.TEXTURE_MAG_FILTER,
		TEXTURE_MAX_ANISOTROPY:                          gl.TEXTURE_MAX_ANISOTROPY,
		TEXTURE_MAX_LEVEL:                               gl.TEXTURE_MAX_LEVEL,
		TEXTURE_MAX_LOD:                                 gl.TEXTURE_MAX_LOD,
		TEXTURE_MIN_FILTER:                              gl.TEXTURE_MIN_FILTER,
		TEXTURE_MIN_LOD:                                 gl.TEXTURE_MIN_LOD,
		TEXTURE_RECTANGLE:                               gl.TEXTURE_RECTANGLE,
		TEXTURE_REDUCTION_MODE_ARB:                      gl.TEXTURE_REDUCTION_MODE_ARB,
		TEXTURE_REDUCTION_MODE_EXT:                      gl.TEXTURE_REDUCTION_MODE_EXT,
		TEXTURE_RED_SIZE:                                gl.TEXTURE_RED_SIZE,
		TEXTURE_RED_TYPE:                                gl.TEXTURE_RED_TYPE,
		TEXTURE_SAMPLES:                                 gl.TEXTURE_SAMPLES,
		TEXTURE_SHADOW:                                  gl.TEXTURE_SHADOW,
		TEXTURE_SHARED_SIZE:                             gl.TEXTURE_SHARED_SIZE,
		TEXTURE_SPARSE_ARB:                              gl.TEXTURE_SPARSE_ARB,
		TEXTURE_SRGB_DECODE_EXT:                         gl.TEXTURE_SRGB_DECODE_EXT,
		TEXTURE_STENCIL_SIZE:                            gl.TEXTURE_STENCIL_SIZE,
		TEXTURE_SWIZZLE_A:                               gl.TEXTURE_SWIZZLE_A,
		TEXTURE_SWIZZLE_B:                               gl.TEXTURE_SWIZZLE_B,
		TEXTURE_SWIZZLE_G:                               gl.TEXTURE_SWIZZLE_G,
		TEXTURE_SWIZZLE_R:                               gl.TEXTURE_SWIZZLE_R,
		TEXTURE_SWIZZLE_RGBA:                            gl.TEXTURE_SWIZZLE_RGBA,
		TEXTURE_TARGET:                                  gl.TEXTURE_TARGET,
		TEXTURE_UPDATE_BARRIER_BIT:                      gl.TEXTURE_UPDATE_BARRIER_BIT,
		TEXTURE_VIEW:                                    gl.TEXTURE_VIEW,
		TEXTURE_VIEW_MIN_LAYER:                          gl.TEXTURE_VIEW_MIN_LAYER,
		TEXTURE_VIEW_MIN_LEVEL:                          gl.TEXTURE_VIEW_MIN_LEVEL,
		TEXTURE_VIEW_NUM_LAYERS:                         gl.TEXTURE_VIEW_NUM_LAYERS,
		TEXTURE_VIEW_NUM_LEVELS:                         gl.TEXTURE_VIEW_NUM_LEVELS,
		TEXTURE_WIDTH:                                   gl.TEXTURE_WIDTH,
		TEXTURE_WRAP_R:                                  gl.TEXTURE_WRAP_R,
		TEXTURE_WRAP_S:                                  gl.TEXTURE_WRAP_S,
		TEXTURE_WRAP_T:                                  gl.TEXTURE_WRAP_T,
		TIMEOUT_EXPIRED:                                 gl.TIMEOUT_EXPIRED,
		TIMEOUT_IGNORED:                                 gl.TIMEOUT_IGNORED,
		TIMESTAMP:                                       gl.TIMESTAMP,
		TIME_ELAPSED:                                    gl.TIME_ELAPSED,
		TOP_LEVEL_ARRAY_SIZE:                            gl.TOP_LEVEL_ARRAY_SIZE,
		TOP_LEVEL_ARRAY_STRIDE:                          gl.TOP_LEVEL_ARRAY_STRIDE,
		TRANSFORM_FEEDBACK:                              gl.TRANSFORM_FEEDBACK,
		TRANSFORM_FEEDBACK_ACTIVE:                       gl.TRANSFORM_FEEDBACK_ACTIVE,
		TRANSFORM_FEEDBACK_BARRIER_BIT:                  gl.TRANSFORM_FEEDBACK_BARRIER_BIT,
		TRANSFORM_FEEDBACK_BINDING:                      gl.TRANSFORM_FEEDBACK_BINDING,
		TRANSFORM_FEEDBACK_BUFFER:                       gl.TRANSFORM_FEEDBACK_BUFFER,
		TRANSFORM_FEEDBACK_BUFFER_ACTIVE:                gl.TRANSFORM_FEEDBACK_BUFFER_ACTIVE,
		TRANSFORM_FEEDBACK_BUFFER_BINDING:               gl.TRANSFORM_FEEDBACK_BUFFER_BINDING,
		TRANSFORM_FEEDBACK_BUFFER_INDEX:                 gl.TRANSFORM_FEEDBACK_BUFFER_INDEX,
		TRANSFORM_FEEDBACK_BUFFER_MODE:                  gl.TRANSFORM_FEEDBACK_BUFFER_MODE,
		TRANSFORM_FEEDBACK_BUFFER_PAUSED:                gl.TRANSFORM_FEEDBACK_BUFFER_PAUSED,
		TRANSFORM_FEEDBACK_BUFFER_SIZE:                  gl.TRANSFORM_FEEDBACK_BUFFER_SIZE,
		TRANSFORM_FEEDBACK_BUFFER_START:                 gl.TRANSFORM_FEEDBACK_BUFFER_START,
		TRANSFORM_FEEDBACK_BUFFER_STRIDE:                gl.TRANSFORM_FEEDBACK_BUFFER_STRIDE,
		TRANSFORM_FEEDBACK_OVERFLOW:                     gl.TRANSFORM_FEEDBACK_OVERFLOW,
		TRANSFORM_FEEDBACK_OVERFLOW_ARB:                 gl.TRANSFORM_FEEDBACK_OVERFLOW_ARB,
		TRANSFORM_FEEDBACK_PAUSED:                       gl.TRANSFORM_FEEDBACK_PAUSED,
		TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:           gl.TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN,
		TRANSFORM_FEEDBACK_STREAM_OVERFLOW:              gl.TRANSFORM_FEEDBACK_STREAM_OVERFLOW,
		TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB:          gl.TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB,
		TRANSFORM_FEEDBACK_VARYING:                      gl.TRANSFORM_FEEDBACK_VARYING,
		TRANSFORM_FEEDBACK_VARYINGS:                     gl.TRANSFORM_FEEDBACK_VARYINGS,
		TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:           gl.TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH,
		TRANSLATE_2D_NV:                                 gl.TRANSLATE_2D_NV,
		TRANSLATE_3D_NV:                                 gl.TRANSLATE_3D_NV,
		TRANSLATE_X_NV:                                  gl.TRANSLATE_X_NV,
		TRANSLATE_Y_NV:                                  gl.TRANSLATE_Y_NV,
		TRANSPOSE_AFFINE_2D_NV:                          gl.TRANSPOSE_AFFINE_2D_NV,
		TRANSPOSE_AFFINE_3D_NV:                          gl.TRANSPOSE_AFFINE_3D_NV,
		TRANSPOSE_PROGRAM_MATRIX_EXT:                    gl.TRANSPOSE_PROGRAM_MATRIX_EXT,
		TRIANGLES:                                       gl.TRIANGLES,
		TRIANGLES_ADJACENCY:                             gl.TRIANGLES_ADJACENCY,
		TRIANGLES_ADJACENCY_ARB:                         gl.TRIANGLES_ADJACENCY_ARB,
		TRIANGLE_FAN:                                    gl.TRIANGLE_FAN,
		TRIANGLE_STRIP:                                  gl.TRIANGLE_STRIP,
		TRIANGLE_STRIP_ADJACENCY:                        gl.TRIANGLE_STRIP_ADJACENCY,
		TRIANGLE_STRIP_ADJACENCY_ARB:                    gl.TRIANGLE_STRIP_ADJACENCY_ARB,
		TRIANGULAR_NV:                                   gl.TRIANGULAR_NV,
		TRUE:                                            gl.TRUE,
		TYPE:                                            gl.TYPE,
		UNCORRELATED_NV:                                 gl.UNCORRELATED_NV,
		UNDEFINED_VERTEX:                                gl.UNDEFINED_VERTEX,
		UNIFORM:                                         gl.UNIFORM,
		UNIFORM_ADDRESS_COMMAND_NV:                      gl.UNIFORM_ADDRESS_COMMAND_NV,
		UNIFORM_ARRAY_STRIDE:                            gl.UNIFORM_ARRAY_STRIDE,
		UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX:             gl.UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX,
		UNIFORM_BARRIER_BIT:                             gl.UNIFORM_BARRIER_BIT,
		UNIFORM_BLOCK:                                   gl.UNIFORM_BLOCK,
		UNIFORM_BLOCK_ACTIVE_UNIFORMS:                   gl.UNIFORM_BLOCK_ACTIVE_UNIFORMS,
		UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES:            gl.UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES,
		UNIFORM_BLOCK_BINDING:                           gl.UNIFORM_BLOCK_BINDING,
		UNIFORM_BLOCK_DATA_SIZE:                         gl.UNIFORM_BLOCK_DATA_SIZE,
		UNIFORM_BLOCK_INDEX:                             gl.UNIFORM_BLOCK_INDEX,
		UNIFORM_BLOCK_NAME_LENGTH:                       gl.UNIFORM_BLOCK_NAME_LENGTH,
		UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER:      gl.UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER,
		UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER:     gl.UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER,
		UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER:     gl.UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER,
		UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER:    gl.UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER,
		UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER: gl.UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER,
		UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:          gl.UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER,
		UNIFORM_BUFFER:                            gl.UNIFORM_BUFFER,
		UNIFORM_BUFFER_ADDRESS_NV:                 gl.UNIFORM_BUFFER_ADDRESS_NV,
		UNIFORM_BUFFER_BINDING:                    gl.UNIFORM_BUFFER_BINDING,
		UNIFORM_BUFFER_LENGTH_NV:                  gl.UNIFORM_BUFFER_LENGTH_NV,
		UNIFORM_BUFFER_OFFSET_ALIGNMENT:           gl.UNIFORM_BUFFER_OFFSET_ALIGNMENT,
		UNIFORM_BUFFER_SIZE:                       gl.UNIFORM_BUFFER_SIZE,
		UNIFORM_BUFFER_START:                      gl.UNIFORM_BUFFER_START,
		UNIFORM_BUFFER_UNIFIED_NV:                 gl.UNIFORM_BUFFER_UNIFIED_NV,
		UNIFORM_IS_ROW_MAJOR:                      gl.UNIFORM_IS_ROW_MAJOR,
		UNIFORM_MATRIX_STRIDE:                     gl.UNIFORM_MATRIX_STRIDE,
		UNIFORM_NAME_LENGTH:                       gl.UNIFORM_NAME_LENGTH,
		UNIFORM_OFFSET:                            gl.UNIFORM_OFFSET,
		UNIFORM_SIZE:                              gl.UNIFORM_SIZE,
		UNIFORM_TYPE:                              gl.UNIFORM_TYPE,
		UNKNOWN_CONTEXT_RESET:                     gl.UNKNOWN_CONTEXT_RESET,
		UNKNOWN_CONTEXT_RESET_ARB:                 gl.UNKNOWN_CONTEXT_RESET_ARB,
		UNKNOWN_CONTEXT_RESET_KHR:                 gl.UNKNOWN_CONTEXT_RESET_KHR,
		UNPACK_ALIGNMENT:                          gl.UNPACK_ALIGNMENT,
		UNPACK_COMPRESSED_BLOCK_DEPTH:             gl.UNPACK_COMPRESSED_BLOCK_DEPTH,
		UNPACK_COMPRESSED_BLOCK_HEIGHT:            gl.UNPACK_COMPRESSED_BLOCK_HEIGHT,
		UNPACK_COMPRESSED_BLOCK_SIZE:              gl.UNPACK_COMPRESSED_BLOCK_SIZE,
		UNPACK_COMPRESSED_BLOCK_WIDTH:             gl.UNPACK_COMPRESSED_BLOCK_WIDTH,
		UNPACK_IMAGE_HEIGHT:                       gl.UNPACK_IMAGE_HEIGHT,
		UNPACK_LSB_FIRST:                          gl.UNPACK_LSB_FIRST,
		UNPACK_ROW_LENGTH:                         gl.UNPACK_ROW_LENGTH,
		UNPACK_SKIP_IMAGES:                        gl.UNPACK_SKIP_IMAGES,
		UNPACK_SKIP_PIXELS:                        gl.UNPACK_SKIP_PIXELS,
		UNPACK_SKIP_ROWS:                          gl.UNPACK_SKIP_ROWS,
		UNPACK_SWAP_BYTES:                         gl.UNPACK_SWAP_BYTES,
		UNSIGNALED:                                gl.UNSIGNALED,
		UNSIGNED_BYTE:                             gl.UNSIGNED_BYTE,
		UNSIGNED_BYTE_2_3_3_REV:                   gl.UNSIGNED_BYTE_2_3_3_REV,
		UNSIGNED_BYTE_3_3_2:                       gl.UNSIGNED_BYTE_3_3_2,
		UNSIGNED_INT:                              gl.UNSIGNED_INT,
		UNSIGNED_INT16_NV:                         gl.UNSIGNED_INT16_NV,
		UNSIGNED_INT16_VEC2_NV:                    gl.UNSIGNED_INT16_VEC2_NV,
		UNSIGNED_INT16_VEC3_NV:                    gl.UNSIGNED_INT16_VEC3_NV,
		UNSIGNED_INT16_VEC4_NV:                    gl.UNSIGNED_INT16_VEC4_NV,
		UNSIGNED_INT64_AMD:                        gl.UNSIGNED_INT64_AMD,
		UNSIGNED_INT64_ARB:                        gl.UNSIGNED_INT64_ARB,
		UNSIGNED_INT64_NV:                         gl.UNSIGNED_INT64_NV,
		UNSIGNED_INT64_VEC2_ARB:                   gl.UNSIGNED_INT64_VEC2_ARB,
		UNSIGNED_INT64_VEC2_NV:                    gl.UNSIGNED_INT64_VEC2_NV,
		UNSIGNED_INT64_VEC3_ARB:                   gl.UNSIGNED_INT64_VEC3_ARB,
		UNSIGNED_INT64_VEC3_NV:                    gl.UNSIGNED_INT64_VEC3_NV,
		UNSIGNED_INT64_VEC4_ARB:                   gl.UNSIGNED_INT64_VEC4_ARB,
		UNSIGNED_INT64_VEC4_NV:                    gl.UNSIGNED_INT64_VEC4_NV,
		UNSIGNED_INT8_NV:                          gl.UNSIGNED_INT8_NV,
		UNSIGNED_INT8_VEC2_NV:                     gl.UNSIGNED_INT8_VEC2_NV,
		UNSIGNED_INT8_VEC3_NV:                     gl.UNSIGNED_INT8_VEC3_NV,
		UNSIGNED_INT8_VEC4_NV:                     gl.UNSIGNED_INT8_VEC4_NV,
		UNSIGNED_INT_10F_11F_11F_REV:              gl.UNSIGNED_INT_10F_11F_11F_REV,
		UNSIGNED_INT_10_10_10_2:                   gl.UNSIGNED_INT_10_10_10_2,
		UNSIGNED_INT_24_8:                         gl.UNSIGNED_INT_24_8,
		UNSIGNED_INT_2_10_10_10_REV:               gl.UNSIGNED_INT_2_10_10_10_REV,
		UNSIGNED_INT_5_9_9_9_REV:                  gl.UNSIGNED_INT_5_9_9_9_REV,
		UNSIGNED_INT_8_8_8_8:                      gl.UNSIGNED_INT_8_8_8_8,
		UNSIGNED_INT_8_8_8_8_REV:                  gl.UNSIGNED_INT_8_8_8_8_REV,
		UNSIGNED_INT_ATOMIC_COUNTER:               gl.UNSIGNED_INT_ATOMIC_COUNTER,
		UNSIGNED_INT_IMAGE_1D:                     gl.UNSIGNED_INT_IMAGE_1D,
		UNSIGNED_INT_IMAGE_1D_ARRAY:               gl.UNSIGNED_INT_IMAGE_1D_ARRAY,
		UNSIGNED_INT_IMAGE_2D:                     gl.UNSIGNED_INT_IMAGE_2D,
		UNSIGNED_INT_IMAGE_2D_ARRAY:               gl.UNSIGNED_INT_IMAGE_2D_ARRAY,
		UNSIGNED_INT_IMAGE_2D_MULTISAMPLE:         gl.UNSIGNED_INT_IMAGE_2D_MULTISAMPLE,
		UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY:   gl.UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY,
		UNSIGNED_INT_IMAGE_2D_RECT:                gl.UNSIGNED_INT_IMAGE_2D_RECT,
		UNSIGNED_INT_IMAGE_3D:                     gl.UNSIGNED_INT_IMAGE_3D,
		UNSIGNED_INT_IMAGE_BUFFER:                 gl.UNSIGNED_INT_IMAGE_BUFFER,
		UNSIGNED_INT_IMAGE_CUBE:                   gl.UNSIGNED_INT_IMAGE_CUBE,
		UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY:         gl.UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY,
		UNSIGNED_INT_SAMPLER_1D:                   gl.UNSIGNED_INT_SAMPLER_1D,
		UNSIGNED_INT_SAMPLER_1D_ARRAY:             gl.UNSIGNED_INT_SAMPLER_1D_ARRAY,
		UNSIGNED_INT_SAMPLER_2D:                   gl.UNSIGNED_INT_SAMPLER_2D,
		UNSIGNED_INT_SAMPLER_2D_ARRAY:             gl.UNSIGNED_INT_SAMPLER_2D_ARRAY,
		UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE:       gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE,
		UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY: gl.UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY,
		UNSIGNED_INT_SAMPLER_2D_RECT:              gl.UNSIGNED_INT_SAMPLER_2D_RECT,
		UNSIGNED_INT_SAMPLER_3D:                   gl.UNSIGNED_INT_SAMPLER_3D,
		UNSIGNED_INT_SAMPLER_BUFFER:               gl.UNSIGNED_INT_SAMPLER_BUFFER,
		UNSIGNED_INT_SAMPLER_CUBE:                 gl.UNSIGNED_INT_SAMPLER_CUBE,
		UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY:       gl.UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY,
		UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB:   gl.UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB,
		UNSIGNED_INT_VEC2:                         gl.UNSIGNED_INT_VEC2,
		UNSIGNED_INT_VEC3:                         gl.UNSIGNED_INT_VEC3,
		UNSIGNED_INT_VEC4:                         gl.UNSIGNED_INT_VEC4,
		UNSIGNED_NORMALIZED:                       gl.UNSIGNED_NORMALIZED,
		UNSIGNED_SHORT:                            gl.UNSIGNED_SHORT,
		UNSIGNED_SHORT_1_5_5_5_REV:                gl.UNSIGNED_SHORT_1_5_5_5_REV,
		UNSIGNED_SHORT_4_4_4_4:                    gl.UNSIGNED_SHORT_4_4_4_4,
		UNSIGNED_SHORT_4_4_4_4_REV:                gl.UNSIGNED_SHORT_4_4_4_4_REV,
		UNSIGNED_SHORT_5_5_5_1:                    gl.UNSIGNED_SHORT_5_5_5_1,
		UNSIGNED_SHORT_5_6_5:                      gl.UNSIGNED_SHORT_5_6_5,
		UNSIGNED_SHORT_5_6_5_REV:                  gl.UNSIGNED_SHORT_5_6_5_REV,
		UNSIGNED_SHORT_8_8_APPLE:                  gl.UNSIGNED_SHORT_8_8_APPLE,
		UNSIGNED_SHORT_8_8_REV_APPLE:              gl.UNSIGNED_SHORT_8_8_REV_APPLE,
		UPPER_LEFT:                                gl.UPPER_LEFT,
		USE_MISSING_GLYPH_NV:                      gl.USE_MISSING_GLYPH_NV,
		UTF16_NV:                                  gl.UTF16_NV,
		UTF8_NV:                                   gl.UTF8_NV,
		VALIDATE_STATUS:                           gl.VALIDATE_STATUS,
		VENDOR:                                    gl.VENDOR,
		VERSION:                                   gl.VERSION,
		VERTEX_ARRAY:                              gl.VERTEX_ARRAY,
		VERTEX_ARRAY_ADDRESS_NV:                   gl.VERTEX_ARRAY_ADDRESS_NV,
		VERTEX_ARRAY_BINDING:                      gl.VERTEX_ARRAY_BINDING,
		VERTEX_ARRAY_KHR:                          gl.VERTEX_ARRAY_KHR,
		VERTEX_ARRAY_LENGTH_NV:                    gl.VERTEX_ARRAY_LENGTH_NV,
		VERTEX_ARRAY_OBJECT_EXT:                   gl.VERTEX_ARRAY_OBJECT_EXT,
		VERTEX_ATTRIB_ARRAY_ADDRESS_NV:            gl.VERTEX_ATTRIB_ARRAY_ADDRESS_NV,
		VERTEX_ATTRIB_ARRAY_BARRIER_BIT:           gl.VERTEX_ATTRIB_ARRAY_BARRIER_BIT,
		VERTEX_ATTRIB_ARRAY_BUFFER_BINDING:        gl.VERTEX_ATTRIB_ARRAY_BUFFER_BINDING,
		VERTEX_ATTRIB_ARRAY_DIVISOR:               gl.VERTEX_ATTRIB_ARRAY_DIVISOR,
		VERTEX_ATTRIB_ARRAY_DIVISOR_ARB:           gl.VERTEX_ATTRIB_ARRAY_DIVISOR_ARB,
		VERTEX_ATTRIB_ARRAY_ENABLED:               gl.VERTEX_ATTRIB_ARRAY_ENABLED,
		VERTEX_ATTRIB_ARRAY_INTEGER:               gl.VERTEX_ATTRIB_ARRAY_INTEGER,
		VERTEX_ATTRIB_ARRAY_LENGTH_NV:             gl.VERTEX_ATTRIB_ARRAY_LENGTH_NV,
		VERTEX_ATTRIB_ARRAY_LONG:                  gl.VERTEX_ATTRIB_ARRAY_LONG,
		VERTEX_ATTRIB_ARRAY_NORMALIZED:            gl.VERTEX_ATTRIB_ARRAY_NORMALIZED,
		VERTEX_ATTRIB_ARRAY_POINTER:               gl.VERTEX_ATTRIB_ARRAY_POINTER,
		VERTEX_ATTRIB_ARRAY_SIZE:                  gl.VERTEX_ATTRIB_ARRAY_SIZE,
		VERTEX_ATTRIB_ARRAY_STRIDE:                gl.VERTEX_ATTRIB_ARRAY_STRIDE,
		VERTEX_ATTRIB_ARRAY_TYPE:                  gl.VERTEX_ATTRIB_ARRAY_TYPE,
		VERTEX_ATTRIB_ARRAY_UNIFIED_NV:            gl.VERTEX_ATTRIB_ARRAY_UNIFIED_NV,
		VERTEX_ATTRIB_BINDING:                     gl.VERTEX_ATTRIB_BINDING,
		VERTEX_ATTRIB_RELATIVE_OFFSET:             gl.VERTEX_ATTRIB_RELATIVE_OFFSET,
		VERTEX_BINDING_BUFFER:                     gl.VERTEX_BINDING_BUFFER,
		VERTEX_BINDING_DIVISOR:                    gl.VERTEX_BINDING_DIVISOR,
		VERTEX_BINDING_OFFSET:                     gl.VERTEX_BINDING_OFFSET,
		VERTEX_BINDING_STRIDE:                     gl.VERTEX_BINDING_STRIDE,
		VERTEX_PROGRAM_POINT_SIZE:                 gl.VERTEX_PROGRAM_POINT_SIZE,
		VERTEX_SHADER:                             gl.VERTEX_SHADER,
		VERTEX_SHADER_BIT:                         gl.VERTEX_SHADER_BIT,
		VERTEX_SHADER_BIT_EXT:                     gl.VERTEX_SHADER_BIT_EXT,
		VERTEX_SHADER_INVOCATIONS:                 gl.VERTEX_SHADER_INVOCATIONS,
		VERTEX_SHADER_INVOCATIONS_ARB:             gl.VERTEX_SHADER_INVOCATIONS_ARB,
		VERTEX_SUBROUTINE:                         gl.VERTEX_SUBROUTINE,
		VERTEX_SUBROUTINE_UNIFORM:                 gl.VERTEX_SUBROUTINE_UNIFORM,
		VERTEX_TEXTURE:                            gl.VERTEX_TEXTURE,
		VERTICAL_LINE_TO_NV:                       gl.VERTICAL_LINE_TO_NV,
		VERTICES_SUBMITTED:                        gl.VERTICES_SUBMITTED,
		VERTICES_SUBMITTED_ARB:                    gl.VERTICES_SUBMITTED_ARB,
		VIEWPORT:                                  gl.VIEWPORT,
		VIEWPORT_BOUNDS_RANGE:                     gl.VIEWPORT_BOUNDS_RANGE,
		VIEWPORT_COMMAND_NV:                       gl.VIEWPORT_COMMAND_NV,
		VIEWPORT_INDEX_PROVOKING_VERTEX:           gl.VIEWPORT_INDEX_PROVOKING_VERTEX,
		VIEWPORT_POSITION_W_SCALE_NV:              gl.VIEWPORT_POSITION_W_SCALE_NV,
		VIEWPORT_POSITION_W_SCALE_X_COEFF_NV:      gl.VIEWPORT_POSITION_W_SCALE_X_COEFF_NV,
		VIEWPORT_POSITION_W_SCALE_Y_COEFF_NV:      gl.VIEWPORT_POSITION_W_SCALE_Y_COEFF_NV,
		VIEWPORT_SUBPIXEL_BITS:                    gl.VIEWPORT_SUBPIXEL_BITS,
		VIEWPORT_SWIZZLE_NEGATIVE_W_NV:            gl.VIEWPORT_SWIZZLE_NEGATIVE_W_NV,
		VIEWPORT_SWIZZLE_NEGATIVE_X_NV:            gl.VIEWPORT_SWIZZLE_NEGATIVE_X_NV,
		VIEWPORT_SWIZZLE_NEGATIVE_Y_NV:            gl.VIEWPORT_SWIZZLE_NEGATIVE_Y_NV,
		VIEWPORT_SWIZZLE_NEGATIVE_Z_NV:            gl.VIEWPORT_SWIZZLE_NEGATIVE_Z_NV,
		VIEWPORT_SWIZZLE_POSITIVE_W_NV:            gl.VIEWPORT_SWIZZLE_POSITIVE_W_NV,
		VIEWPORT_SWIZZLE_POSITIVE_X_NV:            gl.VIEWPORT_SWIZZLE_POSITIVE_X_NV,
		VIEWPORT_SWIZZLE_POSITIVE_Y_NV:            gl.VIEWPORT_SWIZZLE_POSITIVE_Y_NV,
		VIEWPORT_SWIZZLE_POSITIVE_Z_NV:            gl.VIEWPORT_SWIZZLE_POSITIVE_Z_NV,
		VIEWPORT_SWIZZLE_W_NV:                     gl.VIEWPORT_SWIZZLE_W_NV,
		VIEWPORT_SWIZZLE_X_NV:                     gl.VIEWPORT_SWIZZLE_X_NV,
		VIEWPORT_SWIZZLE_Y_NV:                     gl.VIEWPORT_SWIZZLE_Y_NV,
		VIEWPORT_SWIZZLE_Z_NV:                     gl.VIEWPORT_SWIZZLE_Z_NV,
		VIEW_CLASS_128_BITS:                       gl.VIEW_CLASS_128_BITS,
		VIEW_CLASS_16_BITS:                        gl.VIEW_CLASS_16_BITS,
		VIEW_CLASS_24_BITS:                        gl.VIEW_CLASS_24_BITS,
		VIEW_CLASS_32_BITS:                        gl.VIEW_CLASS_32_BITS,
		VIEW_CLASS_48_BITS:                        gl.VIEW_CLASS_48_BITS,
		VIEW_CLASS_64_BITS:                        gl.VIEW_CLASS_64_BITS,
		VIEW_CLASS_8_BITS:                         gl.VIEW_CLASS_8_BITS,
		VIEW_CLASS_96_BITS:                        gl.VIEW_CLASS_96_BITS,
		VIEW_CLASS_BPTC_FLOAT:                     gl.VIEW_CLASS_BPTC_FLOAT,
		VIEW_CLASS_BPTC_UNORM:                     gl.VIEW_CLASS_BPTC_UNORM,
		VIEW_CLASS_RGTC1_RED:                      gl.VIEW_CLASS_RGTC1_RED,
		VIEW_CLASS_RGTC2_RG:                       gl.VIEW_CLASS_RGTC2_RG,
		VIEW_CLASS_S3TC_DXT1_RGB:                  gl.VIEW_CLASS_S3TC_DXT1_RGB,
		VIEW_CLASS_S3TC_DXT1_RGBA:                 gl.VIEW_CLASS_S3TC_DXT1_RGBA,
		VIEW_CLASS_S3TC_DXT3_RGBA:                 gl.VIEW_CLASS_S3TC_DXT3_RGBA,
		VIEW_CLASS_S3TC_DXT5_RGBA:                 gl.VIEW_CLASS_S3TC_DXT5_RGBA,
		VIEW_COMPATIBILITY_CLASS:                  gl.VIEW_COMPATIBILITY_CLASS,
		VIRTUAL_PAGE_SIZE_INDEX_ARB:               gl.VIRTUAL_PAGE_SIZE_INDEX_ARB,
		VIRTUAL_PAGE_SIZE_X_ARB:                   gl.VIRTUAL_PAGE_SIZE_X_ARB,
		VIRTUAL_PAGE_SIZE_Y_ARB:                   gl.VIRTUAL_PAGE_SIZE_Y_ARB,
		VIRTUAL_PAGE_SIZE_Z_ARB:                   gl.VIRTUAL_PAGE_SIZE_Z_ARB,
		VIVIDLIGHT_NV:                             gl.VIVIDLIGHT_NV,
		WAIT_FAILED:                               gl.WAIT_FAILED,
		WARPS_PER_SM_NV:                           gl.WARPS_PER_SM_NV,
		WARP_SIZE_NV:                              gl.WARP_SIZE_NV,
		WEIGHTED_AVERAGE_ARB:                      gl.WEIGHTED_AVERAGE_ARB,
		WEIGHTED_AVERAGE_EXT:                      gl.WEIGHTED_AVERAGE_EXT,
		WINDOW_RECTANGLE_EXT:                      gl.WINDOW_RECTANGLE_EXT,
		WINDOW_RECTANGLE_MODE_EXT:                 gl.WINDOW_RECTANGLE_MODE_EXT,
		WRITE_ONLY:                                gl.WRITE_ONLY,
		XOR:                                       gl.XOR,
		XOR_NV:                                    gl.XOR_NV,
		ZERO:                                      gl.ZERO,
		ZERO_TO_ONE:                               gl.ZERO_TO_ONE,
	}
}

func (c *Context) CreateShader(typ int) *Shader {
	shader := &Shader{gl.CreateShader(uint32(typ))}
	return shader
}

func (c *Context) ShaderSource(shader *Shader, source string) {
	glsource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader.uint32, 1, glsource, nil)
	free()
}

func (c *Context) CompileShader(shader *Shader) {
	gl.CompileShader(shader.uint32)
}

// Ptr takes a slice or pointer (to a singular scalar value or the first
// element of an array or slice) and returns its GL-compatible address.
func (c *Context) Ptr(data interface{}) unsafe.Pointer {
	return gl.Ptr(data)
}

// Str takes a null-terminated Go string and returns its GL-compatible address.
// This function reaches into Go string storage in an unsafe way so the caller
// must ensure the string is not garbage collected.
func (c *Context) Str(str string) *uint8 {
	return gl.Str(str)
}

// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func (c *Context) GoStr(cstr *uint8) string {
	return gl.GoStr(cstr)
}

// DeleteShader will free the shader memory. You should call this in case of
// a compilation error to avoid leaking memory
func (c *Context) DeleteShader(shader *Shader) {
	gl.DeleteShader(shader.uint32)
}

// DeleteTexture will free the texture from the GPU memory
func (c *Context) DeleteTexture(texture *Texture) {
	gl.DeleteTextures(1, &[]uint32{texture.uint32}[0])
}

// Returns a parameter from a shader object
func (c *Context) GetShaderiv(shader *Shader, pname uint32) bool {
	var success int32
	gl.GetShaderiv(shader.uint32, pname, &success)
	return success == int32(gl.TRUE)
}

// GetShaderInfoLog is a method you can call to get the compilation logs of a shader
func (c *Context) GetShaderInfoLog(shader *Shader) string {
	var maxLength int32
	gl.GetShaderiv(shader.uint32, gl.INFO_LOG_LENGTH, &maxLength)

	errorLog := make([]byte, maxLength)
	gl.GetShaderInfoLog(shader.uint32, maxLength, &maxLength, (*uint8)(gl.Ptr(errorLog)))

	return string(errorLog)
}

func (c *Context) CreateProgram() *Program {
	return &Program{gl.CreateProgram()}
}

func (c *Context) DeleteProgram(program *Program) {
	gl.DeleteProgram(program.uint32)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program *Program, index int, name string) {
	gl.BindAttribLocation(program.uint32, uint32(index), gl.Str(name+"\x00"))
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program *Program, pname int) int {
	var success int32 = gl.FALSE
	gl.GetProgramiv(program.uint32, uint32(pname), &success)
	return int(success)
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program *Program, pname int) bool {
	var success int32 = gl.FALSE
	gl.GetProgramiv(program.uint32, uint32(pname), &success)
	return success == gl.TRUE
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program *Program) string {
	var maxLength int32
	gl.GetProgramiv(program.uint32, gl.INFO_LOG_LENGTH, &maxLength)

	errorLog := make([]byte, maxLength)
	gl.GetProgramInfoLog(program.uint32, maxLength, &maxLength, (*uint8)(gl.Ptr(errorLog)))

	return string(errorLog)
}

func (c *Context) AttachShader(program *Program, shader *Shader) {
	gl.AttachShader(program.uint32, shader.uint32)
}

func (c *Context) LineWidth(width float32) {
	gl.LineWidth(width)
}

func (c *Context) LinkProgram(program *Program) {
	gl.LinkProgram(program.uint32)
}

func (c *Context) CreateTexture() *Texture {
	var loc uint32
	gl.GenTextures(1, &loc)
	return &Texture{loc}
}

func (c *Context) BindTexture(target int, texture *Texture) {
	if texture == nil {
		gl.BindTexture(uint32(target), 0)
		return
	}
	gl.BindTexture(uint32(target), texture.uint32)
}

func (c *Context) ActiveTexture(target int) {
	gl.ActiveTexture(uint32(target))
}

func (c *Context) TexParameteri(target int, pname int, param int) {
	gl.TexParameteri(uint32(target), uint32(pname), int32(param))
}

func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, data interface{}) {
	var pix []uint8
	width := 0
	height := 0
	if data == nil {
		pix = nil
	} else {

		switch img := data.(type) {
		case *image.NRGBA:
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			pix = img.Pix
		case *image.RGBA:
			width = img.Bounds().Dx()
			height = img.Bounds().Dy()
			pix = img.Pix
		default:
			panic(fmt.Errorf("Image type unsupported: %T", img))
		}
	}
	gl.TexImage2D(uint32(target), int32(level), int32(internalFormat), int32(width), int32(height), int32(0), uint32(format), uint32(kind), gl.Ptr(pix))
}

func (c *Context) TexImage2DEmpty(target, level, internalFormat, format, kind, width, height int) {
	gl.TexImage2D(uint32(target), int32(level), int32(internalFormat), int32(width), int32(height), int32(0), uint32(format), uint32(kind), nil)
}

func (c *Context) GetAttribLocation(program *Program, name string) int {
	return int(gl.GetAttribLocation(program.uint32, gl.Str(name+"\x00")))
}

func (c *Context) GetUniformLocation(program *Program, name string) *UniformLocation {
	return &UniformLocation{gl.GetUniformLocation(program.uint32, gl.Str(name+"\x00"))}
}

func (c *Context) GetError() int {
	return int(gl.GetError())
}

func (c *Context) CreateBuffer() *Buffer {
	var loc uint32
	gl.GenBuffers(1, &loc)
	return &Buffer{loc}
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer *Buffer) {
	gl.DeleteBuffers(1, &[]uint32{buffer.uint32}[0])
}

func (c *Context) BindBuffer(target int, buffer *Buffer) {
	if buffer == nil {
		gl.BindBuffer(uint32(target), 0)
		return
	}
	gl.BindBuffer(uint32(target), buffer.uint32)
}

func (c *Context) BufferData(target int, data interface{}, usage int) {
	s := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl.BufferData(uint32(target), int(s), gl.Ptr(data), uint32(usage))
}

func (c *Context) EnableVertexAttribArray(index int) {
	gl.EnableVertexAttribArray(uint32(index))
}

func (c *Context) DisableVertexAttribArray(index int) {
	gl.DisableVertexAttribArray(uint32(index))
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	gl.VertexAttribPointer(uint32(index), int32(size), uint32(typ), normal, int32(stride), gl.PtrOffset(offset))
}

func (c *Context) Enable(flag int) {
	gl.Enable(uint32(flag))
}

func (c *Context) IsEnabled(flag int) bool {
	return gl.IsEnabled(uint32(flag))
}

func (c *Context) Disable(flag int) {
	gl.Disable(uint32(flag))
}

func (c *Context) Hint(target, mode int) {
	gl.Hint(uint32(target), uint32(mode))
}

func (c *Context) BlendFunc(src, dst int) {
	gl.BlendFunc(uint32(src), uint32(dst))
}

func (c *Context) BlendEquation(mode int) {
	gl.BlendEquation(uint32(mode))
}

func (c *Context) UniformMatrix2fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix2fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UniformMatrix3fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix3fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UniformMatrix4fv(location *UniformLocation, transpose bool, value []float32) {
	// TODO: count value of 1 is currently hardcoded.
	//       Perhaps it should be len(value) / 16 or something else?
	//       In OpenGL 2.1 it is a manually supplied parameter, but WebGL does not have it.
	//       Not sure if WebGL automatically deduces it and supports count values greater than 1, or if 1 is always assumed.
	gl.UniformMatrix4fv(location.int32, 1, transpose, &value[0])
}

func (c *Context) UseProgram(program *Program) {
	if program == nil {
		gl.UseProgram(0)
		return
	}
	gl.UseProgram(program.uint32)
}

func (c *Context) ValidateProgram(program *Program) {
	if program == nil {
		gl.ValidateProgram(0)
		return
	}
	gl.ValidateProgram(program.uint32)
}

// Specify the value of a uniform variable for the current program object
func (c *Context) Uniform1f(location *UniformLocation, x float32) {
	gl.Uniform1f(location.int32, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location *UniformLocation, x int) {
	gl.Uniform1i(location.int32, int32(x))
}

func (c *Context) Uniform2f(location *UniformLocation, x, y float32) {
	gl.Uniform2f(location.int32, x, y)
}

func (c *Context) Uniform3f(location *UniformLocation, x, y, z float32) {
	gl.Uniform3f(location.int32, x, y, z)
}

func (c *Context) Uniform4f(location *UniformLocation, x, y, z, w float32) {
	gl.Uniform4f(location.int32, x, y, z, w)
}

func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	size := uintptr(reflect.ValueOf(data).Len()) * reflect.TypeOf(data).Elem().Size()
	gl.BufferSubData(uint32(target), offset, int(size), gl.Ptr(data))
}

func (c *Context) DrawArrays(mode, first, count int) {
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
}

func (c *Context) DrawElements(mode, count, typ, offset int) {
	gl.DrawElements(uint32(mode), int32(count), uint32(typ), gl.PtrOffset(offset))
}

func (c *Context) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (c *Context) Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (c *Context) GetViewport() [4]int32 {
	var params [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &params[0])
	return params
}

func (c *Context) Scissor(x, y, width, height int) {
	gl.Scissor(int32(x), int32(y), int32(width), int32(height))
}

func (c *Context) Clear(flags int) {
	gl.Clear(uint32(flags))
}

// CreateRenderBuffer creates a RenderBuffer object.
func (c *Context) CreateRenderBuffer() *RenderBuffer {
	var id uint32
	gl.GenRenderbuffers(1, &id)
	return &RenderBuffer{id}
}

// DeleteRenderBuffer destroys the RenderBufffer object.
func (c *Context) DeleteRenderBuffer(rb *RenderBuffer) {
	gl.DeleteRenderbuffers(1, &rb.uint32)
}

// BindRenderBuffer binds a named renderbuffer object.
func (c *Context) BindRenderBuffer(rb *RenderBuffer) {
	if rb != nil {
		gl.BindRenderbuffer(gl.RENDERBUFFER, rb.uint32)
	} else {
		gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
	}
}

// RenderBufferStorage establishes the data storage, format, and dimensions of a renderbuffer object's image.
func (c *Context) RenderBufferStorage(internalFormat int, width, height int) {
	gl.RenderbufferStorage(gl.RENDERBUFFER, uint32(internalFormat), int32(width), int32(height))
}

// CreateFrameBuffer creates a FrameBuffer object.
func (c *Context) CreateFrameBuffer() *FrameBuffer {
	var id uint32
	gl.GenFramebuffers(1, &id)
	return &FrameBuffer{id}
}

// DeleteFrameBuffer deletes the given framebuffer object.
func (c *Context) DeleteFrameBuffer(fb *FrameBuffer) {
	gl.DeleteFramebuffers(1, &fb.uint32)
}

// BindFrameBuffer binds a framebuffer.
func (c *Context) BindFrameBuffer(fb *FrameBuffer) {
	if fb != nil {
		gl.BindFramebuffer(gl.FRAMEBUFFER, fb.uint32)
	} else {
		gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	}
}

// FrameBufferTexture2D attaches a texture to a FrameBuffer
func (c *Context) FrameBufferTexture2D(target, attachment, texTarget int, t *Texture, level int) {
	gl.FramebufferTexture2D(uint32(target), uint32(attachment), uint32(texTarget), t.uint32, int32(level))
}

// FrameBufferRenderBuffer attaches a RenderBuffer object to a FrameBuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment int, rb *RenderBuffer) {
	gl.FramebufferRenderbuffer(uint32(target), uint32(attachment), gl.RENDERBUFFER, rb.uint32)
}
