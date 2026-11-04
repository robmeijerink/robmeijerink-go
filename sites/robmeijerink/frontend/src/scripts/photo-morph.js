const VERTEX = `
attribute vec2 aPos;
varying vec2 vUv;

void main() {
    vUv = aPos * 0.5 + 0.5;
    gl_Position = vec4(aPos, 0.0, 1.0);
}
`

const FRAGMENT = `
#ifdef GL_FRAGMENT_PRECISION_HIGH
precision highp float;
#else
precision mediump float;
#endif

uniform sampler2D uFrom;
uniform sampler2D uTo;
uniform float uProgress;
uniform float uMaxBlock;
uniform vec2 uSize;
varying vec2 vUv;

float hash(vec2 p) {
    return fract(sin(dot(p, vec2(127.1, 311.7))) * 43758.5453);
}

void main() {
    float bump = pow(sin(3.14159265 * uProgress), 1.4);
    float size = max(1.0, mix(1.0, uMaxBlock, bump));
    float block = floor(exp2(floor(log2(size) * 2.0 + 0.5) / 2.0) + 0.5);

    vec2 px = vUv * uSize;
    vec2 cell = floor(px / block);
    vec2 uv = (cell * block + block * 0.5) / uSize;

    vec2 coarse = floor(px / 14.0);
    float threshold = hash(coarse) * 0.75 + (1.0 - vUv.y) * 0.25;
    float amount = clamp((uProgress - threshold * 0.55) / 0.45, 0.0, 1.0);
    amount = amount * amount * (3.0 - 2.0 * amount);

    gl_FragColor = mix(texture2D(uFrom, uv), texture2D(uTo, uv), amount);
}
`

const DURATION = 600

function ease(value) {
    return value * value * (3 - 2 * value)
}

function compile(gl, type, source) {
    const shader = gl.createShader(type)

    gl.shaderSource(shader, source)
    gl.compileShader(shader)

    return gl.getShaderParameter(shader, gl.COMPILE_STATUS) ? shader : null
}

function createTexture(gl, image, unit) {
    const texture = gl.createTexture()

    gl.activeTexture(gl.TEXTURE0 + unit)
    gl.bindTexture(gl.TEXTURE_2D, texture)
    gl.pixelStorei(gl.UNPACK_FLIP_Y_WEBGL, true)
    gl.pixelStorei(gl.UNPACK_PREMULTIPLY_ALPHA_WEBGL, true)
    gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, image)
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

    return texture
}

function addAlternate(root) {
    const image = document.createElement('img')

    image.src = root.dataset.alt
    image.alt = ''
    image.width = 1000
    image.height = 1144
    image.decoding = 'async'
    image.className = 'hero-photo-alt'
    image.setAttribute('aria-hidden', 'true')
    root.append(image)
}

export async function initPhotoMorph() {
    const root = document.querySelector('.hero-photo')

    if (!root || !root.dataset.alt) return
    if (!window.matchMedia('(hover: hover)').matches) return

    addAlternate(root)

    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return

    const images = root.querySelectorAll('img')

    if (images.length < 2) return

    try {
        await Promise.all(Array.from(images).map((image) => image.decode()))
    } catch {
        return
    }

    const canvas = document.createElement('canvas')
    canvas.className = 'hero-photo-canvas'
    canvas.setAttribute('aria-hidden', 'true')

    const gl = canvas.getContext('webgl', { alpha: true, premultipliedAlpha: true, antialias: false })

    if (!gl) return

    const vertex = compile(gl, gl.VERTEX_SHADER, VERTEX)
    const fragment = compile(gl, gl.FRAGMENT_SHADER, FRAGMENT)

    if (!vertex || !fragment) return

    const program = gl.createProgram()

    gl.attachShader(program, vertex)
    gl.attachShader(program, fragment)
    gl.linkProgram(program)

    if (!gl.getProgramParameter(program, gl.LINK_STATUS)) return

    gl.useProgram(program)

    const buffer = gl.createBuffer()

    gl.bindBuffer(gl.ARRAY_BUFFER, buffer)
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([-1, -1, 1, -1, -1, 1, 1, 1]), gl.STATIC_DRAW)

    const position = gl.getAttribLocation(program, 'aPos')

    gl.enableVertexAttribArray(position)
    gl.vertexAttribPointer(position, 2, gl.FLOAT, false, 0, 0)

    createTexture(gl, images[0], 0)
    createTexture(gl, images[1], 1)

    gl.uniform1i(gl.getUniformLocation(program, 'uFrom'), 0)
    gl.uniform1i(gl.getUniformLocation(program, 'uTo'), 1)

    const progressUniform = gl.getUniformLocation(program, 'uProgress')
    const blockUniform = gl.getUniformLocation(program, 'uMaxBlock')
    const sizeUniform = gl.getUniformLocation(program, 'uSize')

    gl.clearColor(0, 0, 0, 0)

    let progress = 0
    let target = 0
    let last = 0
    let frame = null

    function resize() {
        const rect = root.getBoundingClientRect()
        const ratio = Math.min(window.devicePixelRatio || 1, 2)
        const width = Math.max(1, Math.round(rect.width * ratio))
        const height = Math.max(1, Math.round(rect.height * ratio))

        if (canvas.width !== width || canvas.height !== height) {
            canvas.width = width
            canvas.height = height
            gl.viewport(0, 0, width, height)
        }
    }

    function render() {
        gl.uniform1f(progressUniform, ease(progress))
        gl.uniform1f(blockUniform, canvas.width / 48)
        gl.uniform2f(sizeUniform, canvas.width, canvas.height)
        gl.clear(gl.COLOR_BUFFER_BIT)
        gl.drawArrays(gl.TRIANGLE_STRIP, 0, 4)
    }

    function tick(now) {
        const step = last ? (now - last) / DURATION : 0

        last = now
        progress = target > progress ? Math.min(target, progress + step) : Math.max(target, progress - step)
        render()

        if (progress !== target) {
            frame = requestAnimationFrame(tick)
        } else {
            frame = null
            last = 0
        }
    }

    function go(value) {
        target = value

        if (frame === null) {
            last = 0
            frame = requestAnimationFrame(tick)
        }
    }

    root.append(canvas)
    resize()
    render()
    root.classList.add('is-morph')

    root.addEventListener('pointerenter', (event) => {
        if (event.pointerType === 'mouse') go(1)
    })

    root.addEventListener('pointerleave', (event) => {
        if (event.pointerType === 'mouse') go(0)
    })

    canvas.addEventListener('webglcontextlost', () => {
        root.classList.remove('is-morph')
        canvas.remove()
    })

    function refresh() {
        resize()
        render()
    }

    window.addEventListener('resize', refresh)

    if ('ResizeObserver' in window) {
        new ResizeObserver(refresh).observe(root)
    }
}
