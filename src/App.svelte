<script>
  import svelteLogo from "./assets/svelte.svg";
  import Project from "./lib/Project.svelte";
  // @ts-ignore
  import Banner from "./lib/Banner.svelte";

  let projects = [
    {
      name: "EuRobot Competition",
      stack: [{val:"C/C++",info:"Low level PID controllers, encoders, sensors, etc."},{val:"Matlab",info:'State machine code generation'},
      {val:"Simulink",info:'Modeling and simulating, tuning PID parameters, etc.'},
       {val:"Python",info:"High level control on Raspberry Pi"}],
      desc: "<a target='_blank' href='https://www.eurobot.org/'>EuRobot</a> is an international robotics competition where students from all around the world are building autonomous mobile robots. I joined the Serbian team called <a target='_blank' href='https://robotoid.rs/'>Robotoid</a> in 2018 and stayed until 2021. I was leading the subteam responsible for developing everything related to robot movement and positioning. From the low level PID controllers which regulated left and right motor rpms, tracking position and orientation using wheel encoders, designing and testing Finite-state Machine and communication protocol, up to the Python library which we used to send High-level commands from Raspberry Pi to the robot.  I got very familiar with Arm Cortex-M microcontrollers and our main MCU was STM32F407. I used Matlab & Simulink to make digital models and controllers, which later, using Code Generators from Matlab, I converted to C code and compiled it to run on STM MCU. Because of the sheer complexity starting with Modeling and system identification, PID tuning, embedded system programming, debugging and testing, PCB and hardware manufacturing, etc. I cannot possibly write it all here, but if you are interested for more, in depth, details please feel free to contact me. Regarding EuRobot competition, feel free to contact the Robotoid team or watch a few <a target='_blank' href='https://www.youtube.com/watch?v=lbK0k-ADYrI'>Videos</a> of our and other teams competing and see what it was all about.",
      imgs: ["src/assets/Animation.gif","src/assets/EB7.png", "src/assets/EB1.png","src/assets/EB2.png","src/assets/EB6.png"],
      isOpen:true
    },
    {
      name: "Browser tester",
      stack: [{val:"HTML+JS+CSS",info:'Vanilla page with minimum JS for buttons, popups, etc.'}, {val:"NodeJS",info:"BackEnd code for testing cookies, authorization, etc."}, {val:"GCP",info:"AppEngine hosting, SSL cert, load-balancing"}],
      desc: "This simple website was prof of concept for manual and automatic browser capabilities testing. I developed this app to make browser testing for me and my colleagues easier. After I presented the website to my project lead, we made a completely new browser testing website based on this one and added a lot more capabilities, but it is closed source and used only inside a company.<a target='_blank' href='https://test.lukaluci.com'>Website</a> & <a target='_blank' href='https://github.com/lukadj996/BrowserTest/'>GitHub</a>",
      imgs: ["src/assets/BT1.png","src/assets/BT2.png"],
    },
    {
      name: "Solar energy prject",
      stack: [{val:"GCP",info:"AppEngine hosting, SSL cert, etc."},{val:"Grafana", info:'Data visualization'},{val:"InfluxDB", info:'Time Series database'}, {val:"MCUs",info:'ESP8266 used to read voltage, current, temp, etc.'}],
      desc: "I was always fascinated with renewable energy, especially solar energy. I decided to try and make myself a micro solar plant, just for running some Raspberry Pis as a server, powering security cameras, and other small devices. The goal is to create an off-grid solar-powered server farm for all my future needs, with easy expansion on my mind. One of the requirements I set was to be able to monitor and control the farm remotely. For now, I only have prof of concept done with 4 x 250w panels, connected to few car batteries using Chinese MPPT charging controllers. I was able to power Raspberry Pi, 4G modem, Wi-Fi Cameras and other small devices such as ESP8226, STM32 MCUs for monitoring. Next year I plan to make this at least 3kw with some gel batteries and maybe more durable charging controllers and such.",
      imgs: [
        "src/assets/SL2.png",
        "src/assets/SL1.png",
      ],
    },{
      name: "Betting arbitrage finder",
      stack: [{val:"HTML+JS+CSS",info:'Basic Web stuff, stylng and JS for buttons, popups, etc.'},{val:"jQuery+Datatables",info:'Making sortalbe, searchable, filterable table with grouping, etc.'}, {val:"GCP",info:"AppEngine hosting GoLang backend code, SSL cert, load-balancing"},{val:"SQLite",info:"Simple relational DB for storing data and mapping tables"}],
      desc: "One of many failed projects of mine. But still worth mentioning because I invested enough time and learned new things. The main problem was with scraping all the different betting sites for the same game. I needed to make one-to-many mapping table with unique team names, which ended with over 40k rows in SQL table. Then I scraped as much different betting sites and using GCP Google Search API I linked all the different team names to my unique team name. (for e.g. Man. Utd., Manchester United, Manchester United FC, is the same name for one single team). After that, calculating the best odds and finding arbitrage and how much money to place on each outcome to guarantee mathematical profit was simple. Did it work? Yes. Until it didn't, and I lost real money (I was a student at the time, so even a small amount of money was a lot for me.) At the end of a day, The house always wins, and I paid with my time and money to learn this first hand. Betting companies can, even if you placed your bet, change the odds willingly and can give you a choice of accepting new odd or returning your money, which in this case means losing an arbitrage since you are betting on the same game in 3 different betting companies, and changing or withdrawing from one of them breaks the mathematically guaranteed profit, leaving you with a loss.",
      imgs: ["src/assets/KL1.png","src/assets/KL2.png"],
    },
  ];
</script>

<main>
  <Banner />
  {#each projects as project}
    <Project {project} />
  {/each}
  <div class="made-with">
     <div>Made with Svelte and GCP with &#9825;</div>
    <div>
      <a href="https://svelte.dev" target="_blank" rel="noreferrer"> 
        <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
      </a>
    </div>
  </div>
</main>

<style>
  .made-with{
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
