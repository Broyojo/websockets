import * as THREE from "../lib/three.js-master/build/three.module.js"
import Stats from "../lib/stats.js-master/build/stats.module.js"
import {
    OrbitControls
} from "../lib/three.js-master/examples/jsm/controls/OrbitControls.js"

export function main() {
    const renderer = new THREE.WebGLRenderer({
        antialias: true
    })
    renderer.setSize(window.innerWidth, window.innerHeight);
    document.body.appendChild(renderer.domElement)

    const scene = new THREE.Scene()
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
    const controls = new OrbitControls(camera, renderer.domElement)
    const stats = new Stats()
    stats.showPanel(0)
    document.body.appendChild(stats.dom)

    let total = 0
    const socket = new WebSocket("ws://localhost:8080/ws")
    socket.addEventListener("message", (event) => {
        const box = JSON.parse(event.data)

        const geometry = new THREE.BoxGeometry(box.size, box.size, box.size)
        const material = new THREE.MeshBasicMaterial({
            color: new THREE.Color(box.r, box.g, box.b),
        })
        const mesh = new THREE.Mesh(geometry, material)
        mesh.position.set(box.x, box.y, box.z)
        scene.add(mesh)
        total++
        console.log(total)
    })

    camera.position.z = 2
    controls.update()

    window.addEventListener("resize", () => {
        renderer.setSize(window.innerWidth, window.innerHeight)
        camera.aspect = window.innerWidth / window.innerHeight
        camera.updateProjectionMatrix()
    })

    // render loop
    let update = () => {
        stats.begin()
        requestAnimationFrame(update)
        controls.update()
        renderer.render(scene, camera)
        stats.end()
    }
    update()
}