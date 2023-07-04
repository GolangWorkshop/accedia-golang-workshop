# Homeworks

## 1. Wine club

A captivating new wine store with a geometrical twist has recently opened its doors, intertwining two of our greatest passions: mathematics and fine wine. Within this enigmatic establishment, every wine bottle takes on one of three shapes: cubic, spherical, or pyramidal. While they vary in size, these bottles exclusively adhere to these geometric forms.

Now, Vankata Gopher faces an important task: assisting Accedias wine club in preparing for their upcoming wine tasting event by getting bottles from this exceptional store. Given his responsibility to cater to the entire club, Vankata must calculate the volume of each bottle to ensure there is enough wine for everyone. Moreover, as the bottles need to be elegantly enveloped in branded paper, he also requires knowledge of their surface areas.

Your assistance is vital in helping Vankata fulfill his mission. Write a program capable of performing these essential calculations by providing the dimensions of each bottle. Represent the different bottles via structs and use interfaces where possible. All of the objects should have the .volume() and .surfaceArea() methods. Create a test file for your program and use concurrent testing where possible.With your program, Vankata Gopher can confidently acquire the necessary information, ensuring a splendid wine tasting experience for the Accedias wine club.

### Input

An integer representing the number of bottles of wine. Than for each bottle is given the shape (cube, sphere or pyramid) and its dimensions: the side of the cube, the radius of the sphere, the edge of the pyramid (assume its a tetrahedron). Not all shapes can be present

### Output

The sum of the areas and the sum of the volumes of all wine bottles

<table>
<thead>
  <tr>
    <th>Input</th>
        <th>Processing</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>3 cube 4 sphere 3 pyramid 9</td>
    <td>
    Surface of cube with side 4: 96<br>
    Volume of cube with side 4: 64<br>
    Surface of sphere with radius 3: 113.1<br>
    Volume of sphere with radius 3: 113.1<br>
    Surface of pyramid with edge 9: 140.3<br>
    Volume of pyramid with edge 9: 85.91<br>
    Sum of all surfaces => 96+113.1+140.3 = 349.4 <br>
    Sum of all volumes => 64+113.1+85.91 = 263.01 <br>
    </td>
    <td>
    349.4 263.01
    </td>
  </tr>
    <tr>
    <td>2 cube 10 sphere 13</td>
    <td>
    Surface of cube with side 10: 600<br>
    Volume of cube with side 10: 1000<br>
    Surface of sphere with radius 13: 2123.72<br>
    Volume of sphere with radius 13: 9202.77<br>
    Sum of all surfaces => 600+2123.72 = 2723.72 <br>
    Sum of all volumes => 1000+9202.77 = 10202.77 <br>
    </td>
    <td>
    2723.72 10202.77
    </td>
  </tr>
</tbody>
</table>

## 2. Save the memories

The gophers' last summer vacation was nothing short of amazing, filled with countless photos and videos capturing their unforgettable moments. However, when it comes to uploading these precious memories to the cloud, our team encounters a daunting challenge—they're running out of storage space.

This is where you can make a difference and lend a helping hand. Your task is to implement the Brotli algorithm, a powerful compression and decompression technique, enabling the gophers to optimize their storage and preserve their cherished memories efficiently. Write tests for your program, test whether your brotli compression/decompression works. In the cheatsheet you can find an example of how we can use streams to read a gzipped file from the internet, and then create a stream pipeline that displays info to the user as well as writes the uncompressed data to a file: use it as an example.

### Input

An url to a big file online: For example https://github.com/neo-liang-sap/book/blob/master/Go/The.Go.Programming.Language.pdf

### Output

A compressed with brotli file and then decompressed

## 9. Love story

Gomeo finds himself head over heels in love with CSharpiette, but unfortunately, her family disapproves of their union. Determined to win them over, Gomeo is desperate to make a lasting impression. He seeks your assistance in implementing as many linq functions as you can.

## Notes:

- You can check this repository for ideas https://github.com/ahmetb/go-linq

## 3. Space invaders

The spaceship belonging to our courageous Gopher team has experienced a catastrophic malfunction while on Earth. Following a meticulous investigation, the Gopher team has discovered numerous loose screws as the root cause. Fortunately, you possess a bag of alien wrenches of various sizes, precisely what the Gophers require. Each wrench is assigned a number ranging from 0 to 100, correlating to its respective size. However, not all wrenches hold value in mending the spaceship; only those with prime numbers are required.

Furthermore, to initiate the ship's reactivation, the Gophers need to know the cumulative sum of the wrench sizes utilized during the repair process. The little Gophers have formed two channels: one designated for the acceptance of prime wrenches to facilitate the repairs, while the other serves as a receptacle for discarded wrenches. Seize 20 wrenches at random and dutifully pass them along! Remember to request the sum of these wrenches!

### Input

20 integers from 0 to 100. They can be repetitive

### Output

The sum of the prime numbers and all the discarded wrenches

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Processing</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20</td>
    <td>
    Primes are 2 3 5 7 11 13 17 and 19 : 2 + 3 + 5 + 7 + 11 + 13 + 17 + 19 = 77
    All others are discarded
    </td>
    <td>
    Errors: 1 4 6 8 9 10 12 14 15 16 18 20
    Sum: 77
    </td>
    </tr>
    <tr>
    <td>12 12 83 95 17 84 60 28 51 37 86 52 41 14 93 78 20 90 15 4</td>
    <td>
    Primes are 83 17 41 : 141
    All others are discarded
    </td>
    <td>
    Errors: 12 12 95 84 60 28 51 37 86 52 14 93 78 20 90 15 4
    Sum: 141
    </td>
    </tr>
</tbody>
</table>

## 4. Not so neat list

The Go Workshop is in its preparation phase, and Marto Gopher has made a neat list of website URLs for all the possible prizes for our quiz winners. Happy with his work, he got up to get some coffee, and oops! The Gopher got tangled in his chargers' wires, and the laptop fell dramatically. Good news: Marto and the laptop are fine. Bad news: the neat list is no longer neat. Some of the URLs got messed up and now return errors when called. And worst of all, if he doesn't submit the list by the end of the day, Kiro Gopher won't approve the budget for prizes.

Help our Gopher quickly complete his task. He will give you the list of URLs and expect you to hand him back a map where the key represents the URL and the value represents the response of the URL where no errors were present.

P.S. Due to the super mega giga ultra urgency of the task, use concurrency to fetch the URL content.

### Input

A list of urls to be checked

### Output

A list where next to every non error generating URL is given the response

### Example:

<table>
<thead>
  <tr>
    <th>Input</th>
    <th>Output</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>https://api.chucknorris.io/jokes/random <br>
    https://api.chucknorris.io/jokes/search?query=9 <br>
    https://api.chucknorris.io/jokes/random?category=test<br>
     https://api.chucknorris.io/jokes/random?category=dev<br>
     </td>
     <td>
      https://api.chucknorris.io/jokes/random : 
      {
        "categories": [
            "dev"
        ],
        "created_at": "2020-01-05 13:42:19.104863",
        "icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
        "id": "7ver3y48qqsfktpelir7ua",
        "updated_at": "2020-01-05 13:42:19.104863",
        "url": "https://api.chucknorris.io/jokes/7ver3y48qqsfktpelir7ua",
        "value": "Don't worry about tests, Chuck Norris's test cases cover your code too."
    }<br>
    https://api.chucknorris.io/jokes/random : <br>
    https://api.chucknorris.io/jokes/search?query=9 :{
    "timestamp": "2023-06-21T12:03:51.946Z",
    "status": 400,
    "error": "Bad Request",
    "message": "search.query: size must be between 3 and 120",
    "violations": {
        "search.query": "size must be between 3 and 120"
    }
}<br>
https://api.chucknorris.io/jokes/random?category="dev" : {
      "timestamp": "2023-06-21T12:04:41.919Z",
    "status": 404,
    "error": "Not Found",
    "message": "No jokes for category tets found.",
    "path": "/jokes/random"
}<br>
https://api.chucknorris.io/jokes/random?category=dev : {
    "categories": [
        "dev"
    ],
    "created_at": "2020-01-05 13:42:18.823766",
    "icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
    "id": "hwv5daessqqk0n_sczyima",
    "updated_at": "2020-01-05 13:42:18.823766",
    "url": "https://api.chucknorris.io/jokes/hwv5daessqqk0n_sczyima",
    "value": "No statement can catch the ChuckNorrisException."
}
</td>
</tr>
</tbody>
</table>
