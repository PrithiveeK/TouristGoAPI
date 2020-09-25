package handlers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	Agent "touristapp.com/controllers/agent"
	Models "touristapp.com/models"
)

//GetAllAgents fetches all the agents
func GetAllAgents(c *fiber.Ctx) error {

	agents, err := Agent.GetAll(map[string]int64{
		"type": 2,
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 200,
		"data": agents,
	})
}

//GetAgent fetches the specified agent
func GetAgent(c *fiber.Ctx) error {

	agentID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Printf("Agent id invalid: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	agent, err := Agent.Get(agentID, map[string]int64{
		"type": 2,
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 200,
		"data": agent,
	})
}

//AddAgent adds a new agent
func AddAgent(c *fiber.Ctx) error {
	var newAgent Models.NewAgent

	if err := c.BodyParser(&newAgent); err != nil {
		log.Printf("Cannot parse agent body: %s\n", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}

	agent, err := Agent.Add(&newAgent, 2)
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 500,
		"data": agent,
		"msg":  "Successfully Added!",
	})
}

//GetAllSubAgents fetches all the sub agents
func GetAllSubAgents(c *fiber.Ctx) error {

	agents, err := Agent.GetAll(map[string]int64{
		"type": 3,
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 200,
		"data": agents,
	})
}

//AddSubAgent adds a new agent
func AddSubAgent(c *fiber.Ctx) error {
	var newSubAgent Models.NewAgent

	if err := c.BodyParser(&newSubAgent); err != nil {
		log.Printf("Cannot parse sub agent body: %s\n", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Invalid Data",
		})
	}

	subAgent, err := Agent.Add(&newSubAgent, 3)
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}

	return c.JSON(&fiber.Map{
		"code": 500,
		"data": subAgent,
		"msg":  "Successfully Added!",
	})
}

//GetSubAgent fetches a specified sub agent
func GetSubAgent(c *fiber.Ctx) error {

	subAgentID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Printf("Sub Agent id invalid: %s", err)
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	subAgent, err := Agent.Get(subAgentID, map[string]int64{
		"type": 3,
	})
	if err != nil {
		return c.JSON(&fiber.Map{
			"code": 500,
			"msg":  "Something went wrong. Please, Try again!",
		})
	}
	return c.JSON(&fiber.Map{
		"code": 200,
		"data": subAgent,
	})
}
