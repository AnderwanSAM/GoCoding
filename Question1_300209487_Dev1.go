package main

import (
	"fmt"
	"strconv"
	//"reflect"
)

//Employee

type employee struct {

       staffMember
       socialSecurityNumber string 
       payRate float64
       name  string
       address string
       phone string 
 


}

func (pt employee) toString() string {
               
                  result := "Name: " + pt.name + "\n";
		 result += "Address: " + pt.address + "\n";
		 result += "Phone: " + pt.phone;
		 return result;
}

//cette methode est tres utilisee comme "super" dans les autres classes
 func   newEmployee ( eName string , eAddress string , ePhone string , socSecNumber string , rate float64) *employee{
		 pt := new(employee)
		 pt.name = eName
		 pt.address = eAddress
		 pt.phone = ePhone
		 pt.socialSecurityNumber = socSecNumber
		 pt.payRate = rate
		return pt
}

func (pt employee) pay() float64 {
		return pt.payRate;
}

//Executive

type executive struct{
     employee
     bonus float64
     name  string
     address string
     phone string 
          


}

func (pt executive) toString() string {
               
                  result := "Name: " + pt.name + "\n";
		 result += "Address: " + pt.address + "\n";
		 result += "Phone: " + pt.phone;
		 return result;
}


func  newExecutive( eName string , eAddress string , ePhone string , socSecNumber string , rate float64)  *executive {
                pt := new(executive)
                //pt.newEmployee(eName,eAddress,ePhone,socSecNumber,rate)
                pt.name = eName
		pt.address = eAddress
		pt.phone = ePhone
		pt.socialSecurityNumber = socSecNumber
		pt.payRate = rate
                pt.bonus = 0
                return pt
}


func (pt *executive) awardBonus(execBonus  float64) {
          pt.bonus = execBonus    
}

func (pt executive) pay () float64{
        payment := pt.employee.pay() + pt.bonus 
        pt.bonus = 0
        return payment 
}

//Hourly 

type hourly struct {
       employee
        hoursWorked int
} 

func  newHourly( eName string,eAddress string ,ePhone string,socSecNumber string, rate float64)  *hourly {
            pt := new(hourly)
            pt.hoursWorked = 0
            //pt.newEmployee(eName,eAddress,ePhone,socSecNumber,rate)
                pt.name = eName
		pt.address = eAddress
		pt.phone = ePhone
		pt.socialSecurityNumber = socSecNumber
		pt.payRate = rate
            return pt
}

	
func ( pt *hourly) addHours(moreHours int) {

       pt.hoursWorked += moreHours
}

func( pt hourly) pay() float64 {

         payment := pt.payRate * float64(pt.hoursWorked)
         pt.hoursWorked = 0
         return payment
    
}


func (pt hourly) toString () string {
         result := pt.employee.toString()
         result += " Current hours: "
         result += strconv.Itoa(pt.hoursWorked)
         return result
 

}

//StaffMember

type staffMember interface {
          
	 toString() string
	 pay()  float64
	   

}



//volunteer


type volunteer struct {

                staffMember
                 name  string
                address string
                phone string 
 



}

func (pt volunteer) toString() string {
               
                  result := "Name: " + pt.name + "\n";
		 result += "Address: " + pt.address + "\n";
		 result += "Phone: " + pt.phone;
		 return result;
}


func newVolunteer( eName string, eAddress string, ePhone string) *volunteer{
                 pt := new(volunteer)
                 pt.name = eName
		 pt.address = eAddress
		 pt.phone = ePhone
		 return pt
      
}

func (pt volunteer) pay () float64{ 
           return 0.0
}

//Main 


func main() {
         
  //-----------------------------------------------------------------
 // Sets up the list of staff members.
 //-----------------------------------------------------------------
 
         entry0 := newExecutive ("Sam", "123 Main Line",
		 "555-0469", "123-45-6789", 2423.07)
	 entry1 := newEmployee ("Carla", "456 Off Line",
		 "555-0101", "987-65-4321", 1246.15)
	 entry2 := newEmployee ("Woody", "789 Off Rocker",
		 "555-0000", "010-20-3040", 1169.23)
	 entry3 := newHourly ("Diane", "678 Fifth Ave.",
		 "555-0690", "958-47-3625", 10.55)
	 entry4:= newVolunteer ("Norm", "987 Suds Blvd.",
		 "555-8374")
	 entry5 :=newVolunteer ("Cliff", "321 Duds Lane",
		 "555-7282")
		
		
        staffList :=  [6]staffMember {entry0, entry1, entry2, entry3, entry4,entry5 }
	
	entry0.awardBonus(500.00)
	entry3.addHours(40)
	
	
//-----------------------------------------------------------------
 // Pays all staff members.
 //-----------------------------------------------------------------
                 var amount float64
		
	        for count := 0 ; count < len(staffList) ; count++ {
	                   fmt.Println(staffList[count])
	                   amount = staffList[count].pay()
	                   if (amount == 0.0) {
	                     fmt.Println("Thanks!")
	                    } else {
	                    fmt.Println("Paid: " )  
	                    fmt.Print(amount)
	                    fmt.Println()
	                    }
	
	                    fmt.Println("-----------------------------------")
	        }
	
	
}















