<template>
  <v-container>
  <template>
        <v-data-table
            :headers="headers"
            :items="data.records"
            :items-per-page="15"
            class="elevation-1"
        >
        <template v-slot:item.time="{ item }">
            <span>{{ new Date(item.time).toLocaleString() }}</span>
        </template>
        <template v-slot:item.controls="props">
                <v-btn class="mx-2"  dark  color="primary" @click="onButtonClick(props.item)">
                    <v-icon left> mdi-delete </v-icon>
            Delete
                </v-btn>
            </template>
        </v-data-table>
        <p></p>
  <vue-excel-xlsx
        :data="data.records"
        :columns="columns"
        :filename="'export_data'"
        :sheetname="'sheet 1'"
        > <v-btn class="mx-2" color="primary">
            <v-icon>mdi-download
            </v-icon>
        Download
        </v-btn>
    </vue-excel-xlsx>
    <v-btn class="mx-2" color="primary" @click="calculate()">
        <v-icon left> mdi-align-vertical-bottom</v-icon>
            Evaluate stats
        </v-btn>
        <p></p>

        <div v-if="stats!=''">
            <v-simple-table>
                <template v-slot:default>
                <thead>
                    <tr>
                    <th class="text-left">
                        Stat type
                    </th>
                    <th class="text-left">
                        Value
                    </th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td> CountDifferentInteractions</td>
                        <td>{{ stats.countDifferentInteractions }}</td>
                    </tr>
                    <tr>
                        <td> CalculateTotalTimeOfInteractions</td>
                        <td>{{ stats.calculateTotalTimeOfInteractions }} seconds</td>
                    </tr>
                    <tr>
                        <td> CalculateLongestSequenceOfInput</td>
                        <td>{{ stats.calculateLongestSequenceOfInput }}</td>
                    </tr>
                    <tr>
                        <td> CalculateLongestSequence</td>
                        <td>{{ stats.calculateLongestSequenceType }} : {{ stats.calculateLongestSequenceCounter }}</td>
                    </tr>
                    <tr>
                        <td> CounterTotalInteractions</td>
                        <td>{{ stats.counterTotalInteractions }} </td>
                    </tr>
                    <tr>
                        <td> MaxDelayBetweenInteractions</td>
                        <td>{{ stats.maxDelayBetweenInteractions }} seconds</td>
                    </tr>
                    <tr>
                        <td> MinDelayBetweenInteractions</td>
                        <td>{{ stats.minDelayBetweenInteractions }} seconds</td>
                    </tr>
                    <tr>
                        <td> MeanDelayBetweenInteractions</td>
                        <td>{{ stats.meanDelayBetweenInteractions }} seconds</td>
                    </tr>
                    
                </tbody>
                </template>
            </v-simple-table>
        </div>
        
</template>
  </v-container>
</template>


<script>
import {getFullData, deleteItemWithID, evaluateStats} from '@/api'


  export default {
    data () {
      return {
        data:  '',
        stats: '',
        headers: [
            {
            text: 'ID',
            align: 'start',
            sortable: true,
            value: 'id',
            },
        { text: 'Event type', value: 'eventType' },
        { text: 'HTML tag or URL', value: 'htmlTag' },
        { text: 'Time', value: 'time', dataType: "Date" },
        { text: "Action", value: "controls", sortable: false }
    ],
    columns : [
                {
                    label: "ID",
                    field: "id",
                },
                {
                    label: "Type",
                    field: "eventType",
                },
                {
                    label: "HTML tag or URL",
                    field: "htmlTag",
                },
                {
                    label: "Time",
                    field: "time",
                },
            ],
      }
    },
    mounted() { 
      getFullData().then(response => (this.data = response))
      
  },
  methods: {
      onButtonClick(item){
          console.log("deleting item with ID " + item.id)
          deleteItemWithID(item.id).then(response => (this.data = response))
      },
      calculate(){
            this.stats = evaluateStats().then(response => (this.stats = response))
        }
  },
  
    }
  
</script>