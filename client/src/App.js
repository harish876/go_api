import React,{useState,useEffect} from 'react';
import WebFont from 'webfontloader';
import GoLogo from './img/goLogo.svg';
import DisplayCard from './Components/DisplayCard'
import InputForm from './Components/InputForm';
import { PlusOutlined } from '@ant-design/icons';
import {Layout,theme,Typography,Modal, Button } from 'antd';

//antd design
const { Content } = Layout;
const {Title} = Typography
function App() {
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const showModal = () => {
    setIsModalOpen(true);
  };
  const handleOk = () => {
    setIsModalOpen(false);
  };
  const handleCancel = () => {
    setIsModalOpen(false);
  };
  useEffect(() => {
    WebFont.load({
      google: {
        families: ['Poppins']
      }
    });
   }, []);
  return (
    <Layout>
    <Title  style={{display:"flex",alignItems:'center',textAlign:'center' ,fontFamily: 'Poppins',marginLeft:"40%"}}>
      <img  style={{ position:"relative",marginRight:'10px',width:'40px' ,height:'40x'}}src={GoLogo} alt="Go Icon"/>
        Go Api
    </Title>
    <Button size="middle" style={{width:'5%',marginLeft:'10px',alignItems:'center' ,textAlign:'center' ,backgroundColor:"#4DD0E1" ,color:"white"}} onClick={showModal}>
      <PlusOutlined/>
    </Button>
    <Content style={{padding: '10px 10px',}}>
      <div className="site-layout-content">
        <DisplayCard/>
      </div>
    </Content>
    <Modal title="Create User" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
        <InputForm 
          data={{}}
          type={"create"}
        />
    </Modal>
  </Layout>
  );
}

export default App;
