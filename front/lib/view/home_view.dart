import 'package:flutter/material.dart';
import 'package:front/service/vehicle.dart';
import 'package:front/view/vehicle_view.dart';

class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  final VehicleService _vehicleService = VehicleService();
  var idController = TextEditingController();
  var yearController = TextEditingController();
  var modelController = TextEditingController();
  var plateController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Adicionar novo usuário'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'Id'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: yearController,
                decoration: const InputDecoration(labelText: 'Ano'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: modelController,
                decoration: const InputDecoration(labelText: 'Modelo'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: plateController,
                decoration: const InputDecoration(labelText: 'Placa'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  _vehicleService.postData(
                      name: nameController.text,
                      id: int.tryParse(idController.text) ?? 0);
                  nameController.clear();
                  idController.clear();
                },
                child: const Text('Enviar'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(builder: (context) => const VehiclesView()),
                  );
                },
                child: const Text('Ir para listagem'),
              ),
            )
          ],
        ),
      ),
    );
  }
}
