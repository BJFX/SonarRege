// structure
package SonarRegener

import (
//"fmt"
)

type Bathy struct {
	Header         OIC_Header
	PortAngle      [155]float32
	PortR          [155]float32
	StarboardAngle [300]float32
	StarboardR     [300]float32
}

//initialize bathy struct ,call OIC initialize
func BathyInit() {

}