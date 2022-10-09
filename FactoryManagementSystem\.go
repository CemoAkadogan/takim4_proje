package main

import (
	"fmt"
	"strings"
	"time"
)

const sharp = "#"

type matter struct { // Struct olarak mobilya hammadesi genel tanımlandı
	idNumber  int
	name      string
	stock     float32
	quality   float32
	unitPrice float32
}

type furniture struct { // Struct olarak mobilya genel tanımlandı
	idNumber     int
	typePortable bool
	name         string
	length       float32
	width        float32
	heigth       float32
	color        matter // boyanin ozellikleri
	material     matter // kullanilan malzemenin ozellikleri
	orderNumber  int
}

func progressBar() {
	for i := 0; i < 101; i++ {
		fmt.Printf("\r[%-100s] %d%%", strings.Repeat(sharp, i), i)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("")

}

func factorySize() { // farkli girdirlerde error aliniyor duzenlenebilir, error vermesini veya 0 degeri cikarmasini engelleyerek tekrar kullanicidan veri alinmali
	var factorySizeScan string
	var factorySize int
	var numberOfWorkers int
	fmt.Println("Fabrikanızın boyutununu giriniz.[küçük , orta , büyük ]")
	// Fabrika boyutlari kullanicidan alinirken belli calisan sayisi araliklari ve metre kare degerleri alinip sonra if de islenebilir
	fmt.Scanf("%s", &factorySizeScan)
	if factorySizeScan == "küçük" {
		numberOfWorkers = 5
		factorySize = 300
	} else if factorySizeScan == "orta" {
		numberOfWorkers = 10
		factorySize = 750
	} else if factorySizeScan == "büyük" {
		numberOfWorkers = 15
		factorySize = 1250
	}
	fmt.Println("Fabrikada Çalışan Sayısı ", numberOfWorkers, "\nFabrika Boyutunuz ", factorySize, "m2")

}

func main() {
	// sonsuz dongu

	furnitureS := []furniture{ // kullanici default olarak gelen mobilyalari ozellestirmemis
		// hazir setler kisminda bu yapiyi kullanalim ama bu asamada kullanici mobilya ozelliklerini kendi girdigi halini eklememiz gerekiyor
		{
			name:         "Seat",
			idNumber:     0,
			width:        90,
			heigth:       50,
			length:       110,
			typePortable: false,
		},
		{
			name:         "Chair",
			idNumber:     1,
			width:        35,
			heigth:       96,
			length:       41,
			typePortable: true,
		},
		{
			name:         "Table",
			idNumber:     2,
			width:        80,
			heigth:       75,
			length:       140,
			typePortable: true,
		},
		{
			name:         "Wardrobe",
			idNumber:     3,
			width:        50,
			heigth:       190,
			length:       210,
			typePortable: false,
		},
		{
			name:         "Coffee Table",
			idNumber:     4,
			width:        50,
			heigth:       42,
			length:       50,
			typePortable: true,
		},
		{
			name:         "TV Unit",
			idNumber:     5,
			width:        35,
			heigth:       50,
			length:       180,
			typePortable: false,
		},
	}

	paints := []matter{

		matter{
			idNumber:  1,
			name:      "Best Red Paint",
			stock:     500,
			quality:   9,
			unitPrice: 35,
		},
		matter{
			idNumber:  2,
			name:      "Best Blue Paint",
			stock:     500,
			quality:   9,
			unitPrice: 32,
		},
		matter{
			idNumber:  3,
			name:      "Best Green Paint",
			stock:     500,
			quality:   9,
			unitPrice: 39,
		},
		matter{
			idNumber:  4,
			name:      "Midle Red Paint",
			stock:     500,
			quality:   6,
			unitPrice: 24,
		},
		matter{
			idNumber:  5,
			name:      "Midle Blue Paint",
			stock:     500,
			quality:   6,
			unitPrice: 22,
		},
		matter{
			idNumber:  6,
			name:      "Best Green Paint",
			stock:     500,
			quality:   6,
			unitPrice: 27,
		},
	}
	materials := []matter{

		matter{
			idNumber:  1,
			name:      "Oak Wood",
			stock:     500,
			quality:   9,
			unitPrice: 35,
		},
		matter{
			idNumber:  2,
			name:      "Walnut Wood",
			stock:     500,
			quality:   9,
			unitPrice: 32,
		},
		matter{
			idNumber:  3,
			name:      "Pine Wood",
			stock:     500,
			quality:   9,
			unitPrice: 39,
		},
		matter{
			idNumber:  4,
			name:      "Willow Wood",
			stock:     500,
			quality:   6,
			unitPrice: 24,
		},
		matter{
			idNumber:  5,
			name:      "Cherry Wood",
			stock:     500,
			quality:   6,
			unitPrice: 22,
		},
		matter{
			idNumber:  6,
			name:      "Poplar Wood",
			stock:     500,
			quality:   6,
			unitPrice: 27,
		},
	}
	kitchenSet := []furniture{
		{
			name:         "Table",
			idNumber:     2,
			width:        80,
			heigth:       75,
			length:       140,
			typePortable: true,
		},
		{
			name:         "Chair",
			idNumber:     1,
			width:        35,
			heigth:       96,
			length:       41,
			typePortable: true,
		},
	}
	livingRoomSet := []furniture{
		{
			name:         "Seat",
			idNumber:     0,
			width:        90,
			heigth:       50,
			length:       110,
			typePortable: false,
		},
		{
			name:         "Chair",
			idNumber:     1,
			width:        35,
			heigth:       96,
			length:       41,
			typePortable: true,
		},
		{
			name:         "Table",
			idNumber:     2,
			width:        80,
			heigth:       75,
			length:       140,
			typePortable: true,
		},
		{
			name:         "Coffee Table",
			idNumber:     4,
			width:        50,
			heigth:       42,
			length:       50,
			typePortable: true,
		},
		{
			name:         "TV Unit",
			idNumber:     5,
			width:        35,
			heigth:       50,
			length:       180,
			typePortable: false,
		},
	}
	bedRoomSet := []furniture{
		{
			name:         "Bed",
			idNumber:     6,
			width:        80,
			heigth:       50,
			length:       200,
			typePortable: false,
		},
		{
			name:         "Wardrobe",
			idNumber:     3,
			width:        50,
			heigth:       190,
			length:       210,
			typePortable: false,
		},
		{
			name:         "Nightstand",
			idNumber:     7,
			width:        40,
			heigth:       50,
			length:       40,
			typePortable: true,
		},
	}
	childrensRoom := []furniture{
		{
			name:         "Bed",
			idNumber:     6,
			width:        80,
			heigth:       50,
			length:       200,
			typePortable: false,
		},
		{
			name:         "Wardrobe",
			idNumber:     3,
			width:        50,
			heigth:       190,
			length:       210,
			typePortable: false,
		},
		{
			name:         "Nightstand",
			idNumber:     7,
			width:        40,
			heigth:       50,
			length:       40,
			typePortable: true,
		},
		{
			name:         "Desk",
			idNumber:     8,
			width:        80,
			heigth:       75,
			length:       140,
			typePortable: true,
		},
		{
			name:         "Chair",
			idNumber:     1,
			width:        35,
			heigth:       96,
			length:       41,
			typePortable: true,
		},
	}

	for {
		for {
			var firstChoose string
			var chosenFurniture furniture
			var chosenPaint matter
			var chosenMaterial matter
			var chosenFurnitureID int
			var chosenMaterialID int
			var chosenPaintID int

			fmt.Println("You can choose a single furniture or a set of furnitures :")
			fmt.Println("  ")
			fmt.Println("For single furniture press 'F' , for a set of furnitures press 'S'")
			for (firstChoose != "f") || (firstChoose != "F") || (firstChoose != "S") || (firstChoose != "s") {
				fmt.Scan(&firstChoose)
				if (firstChoose == "f") || (firstChoose == "F") || (firstChoose == "S") || (firstChoose == "s") {
					break
				}

			}

			if (firstChoose == "F") || (firstChoose == "f") {

				furnitureSliceShow(furnitureS)
				fmt.Println("  ")
				fmt.Println("Choose one of this furnitures to customize")
				fmt.Println("  ")
				for {
					fmt.Scan(&chosenFurnitureID)
					if (chosenFurnitureID > len(furnitureS)) || (chosenFurnitureID <= 0) {
						fmt.Println("Please enter a vailed valu for ID number:")

					} else {
						break
					}

				} // mobilya ID secimi icin gereken hatasiz for sonu
				// mobilya secimi ve eslemesi yapildi
				var membersNumber int = 0
				for furnitureS[membersNumber].idNumber != chosenFurnitureID {
					if furnitureS[membersNumber].idNumber == chosenFurnitureID {
						furnitureS[membersNumber] = chosenFurniture

					}
					membersNumber++

				}

				chosenFurniture = furnitureS[membersNumber]
				chosenFurniture.showFurniture()
				// maddesini ve boyasini kullanici girecek
				fmt.Println("  ")

				fmt.Println(" ")
				matterSliceShow(paints)
				fmt.Println("Enter the paint ID number of the paint you want ")
				var mambersofPaintsNumbers int
				// dogru girdi alinacak ve ID paint icin alinmis olacak
				for {
					fmt.Scan(&chosenPaintID)
					if (chosenPaintID > len(paints)) || (chosenPaintID <= 0) {
						fmt.Println("Please enter a true value for ID number:")

					} else {
						break
					}

				}
				// alinan id slice ile karsilastirilip boya tespit edilecek
				for paints[mambersofPaintsNumbers].idNumber != chosenPaintID {
					if paints[mambersofPaintsNumbers].idNumber == chosenPaintID {

					}
					mambersofPaintsNumbers++

				}
				// alinan id secilen boyayla esitlenecek
				chosenPaint = paints[mambersofPaintsNumbers]
				chosenFurniture.color = chosenPaint
				fmt.Println(" ")
				fmt.Println(" ")
				// materyal secimi yaptirilacak
				fmt.Println("***** MATERIALS*****")
				fmt.Println(" ")

				var memberNumberofMaterials int

				fmt.Println(" ")
				matterSliceShow(materials)
				fmt.Println(" ")
				fmt.Println("Chose one of the materials you want :")
				for {
					fmt.Scan(&chosenMaterialID)
					if (chosenMaterialID > len(materials)) || (chosenMaterialID <= 0) {
						fmt.Println("Please enter a true valu for ID number:")

					} else {
						break
					}

				}
				for materials[memberNumberofMaterials].idNumber != chosenMaterialID {
					if materials[memberNumberofMaterials].idNumber == chosenMaterialID {

					}

					memberNumberofMaterials++

				}
				chosenMaterial = materials[memberNumberofMaterials]
				chosenFurniture.material = chosenMaterial
				fmt.Println("  ")
				fmt.Println("  ")
				fmt.Println("Furniture you have costumized :")
				chosenFurniture.showFurniture()
				fmt.Println("The price of the ", chosenFurniture.name, "is :", priceCalculator(chosenFurniture))
				orderProgres(chosenFurniture)
				fmt.Println(chosenFurniture.material.stock)
				fmt.Println(chosenFurniture.color.stock)
			} // Furniture sectiyse burasai calisacak

		}

	}
	// ctrl c harici calismaya devam etmesi icin

}

func (furn furniture) showFurniture() {
	fmt.Println("Furniture name  : ", furn.name)
	if furn.typePortable == true {
		fmt.Println(furn.name, "is portable")
		fmt.Println("  ")

	} else {
		fmt.Println(furn.name, "is not portable")
		fmt.Println("  ")
	}
	fmt.Println("ID number of", furn.name, "is : ", furn.idNumber)
	fmt.Println("  ")
	if furn.length != 0 {
		fmt.Println("Length of", furn.name, "is : ", furn.length)
		fmt.Println("  ")
	}
	if furn.width != 0 {
		fmt.Println("Width of", furn.name, "is : ", furn.width)
		fmt.Println("  ")
	}
	if furn.heigth != 0 {
		fmt.Println("Higth of", furn.name, "is : ", furn.width)
		fmt.Println("  ")
	}
	if furn.color.idNumber > 0 {
		fmt.Print("The color you chose is : ")
		furn.color.showMaterial()
		fmt.Println("  ")
	}
	if furn.material.idNumber > 0 {
		fmt.Print("Material you chose is")
		furn.material.showMaterial()

	}

}

func (raw matter) showMaterial() {
	fmt.Println(raw.name)
	fmt.Println("ID number of", raw.name, "is : ", raw.idNumber)
	fmt.Println("Stock amount : ", raw.stock)
	fmt.Println("Quality value out of 10 is : ", raw.quality)
	fmt.Println("Unit Price : ", raw.unitPrice, "TL")

}
func furnitureSliceShow(furnitureSet []furniture) { // slice elemanlarını sergilememizi sağlar
	fmt.Println("*****  FURNITURES  *****")
	fmt.Println(" ")
	var i int = len(furnitureSet) - 1
	for 0 <= i {
		furnitureSet[i].showFurniture()
		fmt.Println(" ")

		i -= 1
	}
}
func matterSliceShow(matters []matter) {
	fmt.Println("*****  Matter  *****")
	fmt.Println(" ")
	var n int = len(matters) - 1
	for 0 <= n {
		matters[n].showMaterial()
		fmt.Println(" ")

		n -= 1
	}

}
func orderProgres(furn furniture) {
	var neededMaterial float32
	neededMaterial = furn.heigth * furn.length * furn.width
	if (furn.material.stock < neededMaterial) || (furn.color.stock < neededMaterial) {
		fmt.Println("There is not enough material for production ")
		fmt.Println("Material order is preparing")
		furn.color.stock = neededMaterial
		furn.material.stock = neededMaterial
		progressBar()
		fmt.Println("Stock has been updated. ")

	}

	if (furn.material.stock >= neededMaterial) || (furn.color.stock >= neededMaterial) {
		progressBar()
		var newStockMaterial float32
		var newStockColor float32
		newStockColor = furn.color.stock - neededMaterial
		newStockMaterial = furn.material.stock - neededMaterial
		furn.color.stock = newStockColor
		furn.material.stock = newStockMaterial

	}

}
func priceCalculator(furn furniture) float32 {
	price := (furn.length * furn.heigth * furn.width) * (furn.color.unitPrice) * (furn.material.unitPrice) / 100

	return price
}
