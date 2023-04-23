import React, {useState, useEffect} from 'react';
import './App.css';
import {Greet, CPU, DISK} from "../wailsjs/go/main/App";
import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend, CategoryScale, registerables } from 'chart.js';

ChartJS.register(ArcElement, Tooltip, Legend);

function App() {


    // ----------- CPU
    const [cpu, setcpu] = useState('');
    const updateResultText = (result) => setcpu(result);

    useEffect(() => {

        setInterval(() => {
            CPU().then(updateResultText)
        }, 2000);
    })

    // ------------ RAM
    const [disk, setdisk] = useState('');
    const updateResultDisk = (result) => setdisk(result);

    useEffect(() => {

        setInterval(() => {
            DISK().then(updateResultDisk)
        }, 5000);
    })

    const data = {
        labels: ["CPU", "NO CPU"],
        datasets: [
            {
                label: '% del CPU',
                data: [cpu, (100-cpu)],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)',
                ],
                borderWidth: 1,
            },
        ]
    };

    var optionsCpu = {
        responsive: true,
        maintainAspectRatio: false,
        title: {
            display: true,
            text: 'Gráfica de Pie'
        },
        legend: {
            display: true,
            position: 'bottom',
            labels: {
                fontColor: '#333',
                fontSize: 12
            }
        }
    };


    const dataDisk = {
        labels: ["RAM", "NO RAM"],
        datasets: [
            {
                label: '% de RAM',
                data: [disk, (100-disk)],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)',
                ],
                borderWidth: 1,
            },
        ],
    };



    return (
        <div >

            <br />
            <br />
            <br />

            <div id="title">
                Grafica de PIE, Uso de RAM
            </div>

            <div >
                <Pie
                    data={dataDisk}
                    options={{
                        responsive: true,
                        maintainAspectRatio: false, // Agrega esta opción para deshabilitar la relación de aspecto predeterminada
                        width: 500, // Especifica el ancho de la gráfica
                        height: 500, // Especifica la altura de la gráfica
                        title:{
                            display:true,
                            text:'Porcentaje de la RAM',
                            fontSize:20
                        },
                        legend:{
                            display:true,
                            position:'right'
                        }
                    }}
                />
            </div>
            <div>
                RAM =  %{disk}
            </div>
        </div>
    )
}

export default App
